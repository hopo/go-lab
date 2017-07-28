package main03

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":4000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, `<h1>Hello Wrold</h1>`)
}
