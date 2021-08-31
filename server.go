package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
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
	log.Println(r.URL.Path)
	name := strings.TrimLeft(r.URL.Path, "/candidate/add/")
	log.Println(name)
	for _, val := range candidates {
		if val.name == name {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Candidate already exists")
			return
		}
	}
	candidates = append(candidates, candidate{name, 0})
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, fmt.Sprintf("Candidate %q added", name))
}
func deleteCandidate(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimLeft(r.URL.Path, "/candidate/delete/")
	fmt.Fprintf(w, fmt.Sprintf("%q", name))
}
func voteCandidate(w http.ResponseWriter, r *http.Request) {
	// name := strings.Trim(r.URL.Path, "/candidate/vote/")
}
func votingStatus(w http.ResponseWriter, r *http.Request) {
}
func votingResult(w http.ResponseWriter, r *http.Request) {
}
