package main

import (
	"fmt"
	"log"
	"net/http"

	"app/templates"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.Index().Render(r.Context(), w)
	})

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from the updated server!")
	})

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
