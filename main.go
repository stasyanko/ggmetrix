package main

import (
	"net/http"

	"github.com/gobuffalo/packr"
)

// building GOPATH/bin/packr main.go
func main() {
	box := packr.NewBox("./static")

	http.Handle("/", http.FileServer(box))
	http.ListenAndServe(":3000", nil)
}
