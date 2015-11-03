package main

import (
	"io"
	"log"
	"net/http"
)

const baseGoURL = "https://go.googlesource.com/go/+/"
const version = "go1.5"

func checkURL(url string) bool {
	r, err := http.Head(url)
	if err != nil {
		log.Print(err)
		return false
	}
	return r.StatusCode == http.StatusOK
}

func handleHttp(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, version + " is ")
	if checkURL(baseGoURL + version) {
		io.WriteString(w, "out!")
		return
	}
	io.WriteString(w, "not out yet...")
}

func main() {
	http.HandleFunc("/", handleHttp)
	http.ListenAndServe(":8080", nil)
}
