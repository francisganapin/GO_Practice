package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body><h1>hello,World<h1></body></html>")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
