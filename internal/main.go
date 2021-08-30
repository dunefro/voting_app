package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, fmt.Sprintf("Hello, %q", html.EscapeString(r.URL.Path)))
	})

	log.Println("Listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
