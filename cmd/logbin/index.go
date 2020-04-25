package main

var indexPage = `<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>logbin</title>

    <style type="text/css">
        html, body, div, span, applet, object, iframe,
        h1, h2, h3, h4, h5, h6, p, blockquote, pre,
        a, abbr, acronym, address, big, cite, code,
        del, dfn, em, img, ins, kbd, q, s, samp,
        small, strike, strong, sub, sup, tt, var,
        b, u, i, center,
        dl, dt, dd, ol, ul, li,
        fieldset, form, label, legend,
        table, caption, tbody, tfoot, thead, tr, th, td,
        article, aside, canvas, details, embed,
        figure, figcaption, footer, header, hgroup,
        menu, nav, output, ruby, section, summary,
        time, mark, audio, video {
            margin: 0;
            padding: 0;
            border: 0;
            font-size: 100%;
            font: inherit;
            vertical-align: baseline;
        }

        article, aside, details, figcaption, figure,
        footer, header, hgroup, menu, nav, section {
            display: block;
        }

        body {
            line-height: 1;
        }

        ol, ul {
            list-style: none;
        }

        blockquote, q {
            quotes: none;
        }

        blockquote:before, blockquote:after,
        q:before, q:after {
            content: '';
            content: none;
        }

        table {
            border-collapse: collapse;
            border-spacing: 0;
        }

        body {
            font-size: 14px;
            line-height: 1.25em;
            background-color: #f0f0f0;
        }

        .wrapper {
            display: flex;
            flex-direction: row;
            justify-content: center;
            width: 100%;
        }

        .column {
            width: 70em;
            padding: 1rem 0 1rem 0;
        }

        .page {
            background-color: #fefefe;
            padding: 0.5rem 2rem 3rem 2rem;
        }

        .header {
            font-size: 1.2em;
            font-family: Consolas, monospace;
            margin-top: 1rem;
            padding: 0.5em 0 0.5em 0;
        }

        .header a {
            text-decoration: none;
        }

        .header a:hover {
            text-decoration: underline;
        }

        .header span.red {
            color: #b30014;
        }

        .header span.part {
            color: #666;
            padding-left: 0.2em;
        }

        .header span.part a {
            color: rgb(27, 106, 203);
        }

        .header span.part a:visited {
            color: rgb(27, 106, 203);
        }

        .footer {
            font-size: 0.8em;
            color: #ccc;
            font-weight: 800;
            font-family: helvetica, arial, sans-serif;
            padding: 0.5em 1em 1em;
            text-align: right;
        }

        .footer .left {
            float: left;
        }

        .footer .right {
            float: right;
        }

        .footer a {
            color: #bbb;
        }

        h1, h2, h3, h4 {
            font-family: helvetica, arial, sans-serif;
        }

        .content h1 {
            font-size: 1.6em;
            padding: 1em 0 0 0;
            font-weight: 800;
        }

        .content h2 {
            font-size: 1.3em;
            padding: 0.8em 0 0 0;
            color: #333;
            font-weight: 800;
        }

        .content h3 {
            font-size: 1.2em;
            padding: 0.4em 0 0 0;
            color: #444;
        }

        .content h4 {
            font-size: 1.0em;
            color: #555;
        }

        .content strong {
            font-weight: 600;
        }

        .content code {
            font-family: Consolas, monospace;
            background-color: #f8f8f8;
        }

        .content pre {
            background-color: #f8f8f8;
            border: 1px solid #d8d8d8;
            margin: 1em;
            padding: 0.5em;
            overflow: auto;
        }

        .content p {
            margin-top: 0.8em;
            line-height: 1.5em;
        }

        .content ul {
            padding-top: 0.5em;
            line-height: 1.5em;
        }

        .content ul li {
            padding-left: 1em;
        }

        .content ul li::before {
            content: "•";
            color: #333;;
            display: inline-block;
            width: 1em;
            margin-left: -1em;
        }

        .content img {
            max-width: 90%;
            margin: 1em auto 1em auto;
            display: block;
        }

        .toc {
            float: right;
            padding: 1em 1em 1em 1em;
            border: 1px solid #ddd;
            background-color: #f8f8f8;
            margin: 2em;
            max-width: 30%;
        }

        .toc h1 {
            font-size: 1.2em;
            padding-bottom: 0.5em;
        }

        .toc a {
            text-decoration: none;
        }

        .toc li {
            padding-left: 0.5em;
        }

        .toc ul {
            list-style-type: disc;
            padding-left: 1em;
        }

        .toc ul ul {
            list-style-type: circle;
        }

    </style>
</head>
<body>
<div class="wrapper">
    <div class="column">
        <div class="page">
            <div class="content">
                <h1 id="toc_0">logbin</h1>
				<p>
					Logs can be uploaded as a a POST request, with the log data as plain request body.
					The path of your request will be the name of your upload. It should contain an
					appropriate file ending (.log or .txt for plain text, .xz or .gz for compressed files).
				</p>
                <p>Example using <code>curl</code>:</p>
				<code>curl --data-binary @logfile.txt '{{.publicURL}}'</code>

				<p>journalctl example:</p>
				<code>journalctl -u name.service -o cat --utc --since '09:30' --until '10:30' | curl --data-binary @- '{{.publicURL}}'</code>

				<p>
					If your upload has successfully completed, the server will reply with a confirmation message.
					Uploads are not publicly accessible.
				</p>

				<p>The upload size limit is {{.limitMB}} MB.</p>
				
				<p>If you need to upload a large log file (>100MB), please compress it before uploading it:</p>
				<code>xz < logfile.txt | curl --data-binary @- '{{.publicURL}}.xz'</code>
            </div>
        </div>
        <div class="footer">
            <div class="right">Powered by <a href="https://github.com/leoluk/logbin" target="_blank">logbin</a>.</div>
        </div>
    </div>
</div>
</body>
</html>`
