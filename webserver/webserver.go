package main

import (
	"fmt"
	"net/http"
)

var port int = 12345;

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Printf("Opened web server on port: %d\n", port)
    http.ListenAndServe(":12345", nil)
}