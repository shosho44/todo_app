package main

import (
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/index.html")
	t.Execute(w, nil)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
