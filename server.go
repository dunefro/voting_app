package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/healthz", healthz)

	log.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Up and running !")
}
