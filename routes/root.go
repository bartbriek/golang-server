package routes

import (
	"fmt"
	"io"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received a request on: /\n")
	io.WriteString(w, "<html><body><h1>This is my webiste!</h1></body></html>")
}
