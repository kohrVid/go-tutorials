package main

import (
	"net/http"
	"text/template"
)

const header = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Example Title</title>
</head>
`

const footer = `
</html>
`

const doc = `
{{template "header" .Title}}
<body>
  <h1>Hello {{.Title}}!</h1>
  {{if eq .Path "/Google"}}
    <h2>Hey, Google made Go!</h2>
  {{else}}
    <h2>Hello, {{.Path}}</h2>
  {{end}}

  <ul>
    {{range .Fruits}}
      <li>{{.}}</li>
    {{end}}
  </ul>
</body>
{{template "footer"}}
`

type Context struct {
	Fruits [3]string
	Title  string
	Path   string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		templates := template.New("template")
		templates.New("test").Parse(doc)
		templates.New("header").Parse(header)
		templates.New("footer").Parse(footer)
		context := Context{
			[3]string{"Lemon", "Orange", "Apple"},
			"the title",
			req.URL.Path,
		}
		templates.Lookup("test").Execute(w, context)
	})
	http.ListenAndServe(":9000", nil)
}
