package main

import (
	"io"
	"net/http"
)

func handleHttp(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func main() {
	http.HandleFunc("/", handleHttp)
	http.ListenAndServe(":8080", nil)
}
