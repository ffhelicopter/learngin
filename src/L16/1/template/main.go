package main

import (
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("param")
	tmpl := template.New("hello")
	tmpl, _ = tmpl.Parse(`{{define "T"}}{{.}}{{end}}`)
	tmpl.ExecuteTemplate(w, "T", param1)
}

func main() {
	http.HandleFunc("/test", handler)
	http.ListenAndServe(":8080", nil)
}
