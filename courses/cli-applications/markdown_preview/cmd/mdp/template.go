package main

import "html/template"

type templateParams struct {
	Title string
	Body  template.HTML
}

const (
	defaultTemplate = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="content-type" content="text/html; charset=utf-8">
<title>{{ .Title }}</title>
</head>
<body>
{{ .Body }}
</body>
</html>
`
)
