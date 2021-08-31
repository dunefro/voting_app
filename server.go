package main

import (
	"fmt"
	"log"
	"net/http"
)

type candidate struct {
	name  string
	votes int
}

var candidates []candidate

func main() {
	// http.HandleFunc("/", root)
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/candidate/list/", listCandidate)
	http.HandleFunc("/candidate/add/", addCandidate)
	http.HandleFunc("/candidate/delete/", deleteCandidate)
	http.HandleFunc("/candidate/vote/", voteCandidate)
	http.HandleFunc("/voting/status", votingStatus)
	http.HandleFunc("/voting/results", votingResult)
	log.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Up and running !")
}
func listCandidate(w http.ResponseWriter, r *http.Request) {
	ListOfCandidates := fmt.Sprintf("List of candidates -\n%+v", candidates)
	fmt.Fprintf(w, ListOfCandidates)
}
func addCandidate(w http.ResponseWriter, r *http.Request) {
	// name := strings.Trim(r.URL.Path, "/candidate/add/")
}
func deleteCandidate(w http.ResponseWriter, r *http.Request) {
	// name := strings.Trim(r.URL.Path, "/candidate/delete/")
}
func voteCandidate(w http.ResponseWriter, r *http.Request) {
	// name := strings.Trim(r.URL.Path, "/candidate/vote/")
}
func votingStatus(w http.ResponseWriter, r *http.Request) {
}
func votingResult(w http.ResponseWriter, r *http.Request) {
}
