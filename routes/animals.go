package routes

import (
	"fmt"
	"io"
	"net/http"
)

// List all animals in database
func GetAnimals(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received a request on: /animals\n")
	io.WriteString(w, "{\"value\": \"Animals\": [\"Lion\"]}")
}

// List details of a specific animal
func GetAnimalById(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received a request on: /animals/{{id}}\n")
}
