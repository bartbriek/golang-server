package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	routes "golang-server/routes"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.GetRoot)
	mux.HandleFunc("/animals", routes.GetAnimals)

	err := http.ListenAndServe(":3333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server was closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
