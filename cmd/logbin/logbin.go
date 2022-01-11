package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
)

var (
	token     = flag.String("token", "public", "Subdirectory for directory listing")
	dataDir   = flag.String("dataDir", "", "Data directory (required)")
	addr      = flag.String("addr", ":8080", "Listen address")
	publicURL = flag.String("publicURL", "", "Domain that is displayed on the index page")
	limit     = flag.Int64("limit", 1000000000, "Upload size limit") // 1 GB
)

var indexTmpl = template.Must(template.New("index").Parse(indexPage))

func init() {
	flag.Parse()

	if *dataDir == "" {
		log.Fatal("No data directory specified")
	}

	if *publicURL == "" {
		log.Fatal("No public URL specified")
	}
}

func main() {
	fs := http.FileServer(http.Dir(*dataDir))
	http.Handle(fmt.Sprintf("/%s/", *token), http.StripPrefix("/"+*token, fs))

	u, err := url.Parse(*publicURL)
	if err != nil {
		panic(err)
	}
	u.Path = path.Join(u.Path, "/path/of/your/upload.log")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s %s %s", r.RemoteAddr, r.Method, r.URL, r.Proto, r.Header)

		switch r.Method {
		case "GET":
			err := indexTmpl.Execute(w, map[string]string{
				"publicURL": u.String(),
				"limitMB":   strconv.Itoa(int(*limit) / 1000000),
			})
			if err != nil {
				log.Printf("[%s] error writing response: %s", r.RemoteAddr, err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "POST":
			clean := path.Clean(r.URL.Path)
			p := path.Join(*dataDir, clean)
			dir, _ := path.Split(p)

			err := os.MkdirAll(dir, 0770)
			if err != nil {
				log.Printf("[%s] error creating directory: %w", r.RemoteAddr, err)
				http.Error(w, "Error creating directory", 500)
				return
			}

			f, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
			if os.IsExist(err) {
				log.Printf("[%s] file exists", r.RemoteAddr)
				http.Error(w, "Existing upload at this location", 409)
				return
			} else if err != nil {
				log.Printf("[%s] error creating file: %v", r.RemoteAddr, err)
				http.Error(w, "Error creating file", 500)
				return
			}

			defer f.Close()

			n, err := io.CopyN(f, r.Body, *limit)
			if err != nil && err != io.EOF {
				log.Printf("[%s] %v", r.RemoteAddr, err)
				http.Error(w, fmt.Sprintf("Server Error (EOF?)", *limit), 500)
				return
			} else if err != io.EOF && n == *limit {
				log.Printf("[%s] upload truncated", r.RemoteAddr)
				http.Error(w, fmt.Sprintf("Your upload is too large and has been truncated to %d bytes", *limit), 413)
				return
			}

			if n <= 32 {
				log.Printf("[%s] upload too small (%d bytes)", r.RemoteAddr, n)
				http.Error(w, fmt.Sprintf("You uploaded only %d bytes - please check your command and try again (minimum is 32)", n), 400)
				os.Remove(p)
				return
			}

			w.WriteHeader(200)
			w.Write([]byte(fmt.Sprintf("Uploaded %d bytes to %s\n", n, clean)))
			log.Printf("[%s] completed upload to %s", r.RemoteAddr, p)
		}

	})

	log.Printf("Listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
