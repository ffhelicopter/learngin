package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get("param"))
}

func main() {
	http.HandleFunc("/test", handler)
	http.ListenAndServe(":8080", nil)
}
