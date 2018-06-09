package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.Response, req *http.Request) {
	switch req.URL.Path {
	case "/cat":
		io.WriteString(res, "kitty kitty kitty")
	case "/dog":
		io.WriteString(res, "doggy doggy doggy")
	}
}

func main() {
	var h hotdog
	http.ListenAndServe(":4000", h)
}
