package main

import (
	"io"
	"log"
	"net/http"
)

func handleHttp(w http.ResponseWriter, r *http.Request) {
	ret, err := http.Head("https://go.googlesource.com/go/+/go1.6")
	if err != nil {
		log.Print(err)
		io.WriteString(w, "error from http.Head()")
	}
	if ret.StatusCode != http.StatusOK {
		io.WriteString(w, "Go 1.6 is not out yet...")
		return
	}
	io.WriteString(w, "Go 1.6 is out!")
}

func main() {
	http.HandleFunc("/", handleHttp)
	http.ListenAndServe(":8080", nil)
}
