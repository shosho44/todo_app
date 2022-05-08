package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/index.html")
	err := t.Execute(w, nil)
	if err != nil {
		fmt.Println("error occured inner handler")
	}
}
func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("error occured inner main")
	}
}
