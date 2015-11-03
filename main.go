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

type Server struct {
	version string
	url string
	out bool
}

func NewServer(version, url string) *Server {
	s := &Server{version: version, url: url, out: false}
	s.check()
	return s
}

func (s *Server) check() {
	if !s.out {
		s.out = checkURL(s.url + s.version)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, version + " is ")
	if s.out {
		io.WriteString(w, "out!")
		return
	}
	io.WriteString(w, "not out yet...")
}

func main() {
	http.Handle("/", NewServer(version, baseGoURL))
	http.ListenAndServe(":8080", nil)
}
