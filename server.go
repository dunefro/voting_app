package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type candidate struct {
	name  string
	votes int
}

var candidates []*candidate

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
	var sb strings.Builder
	for _, candidate := range candidates {
		sb.WriteString(candidate.name)
	}
	fmt.Fprintf(w, sb.String())
}
func addCandidate(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/candidate/add/")
	for _, val := range candidates {
		if val.name == name {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Candidate already exists")
			return
		}
	}
	candidates = append(candidates, &candidate{name, 0})
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, fmt.Sprintf("Candidate %q added", name))
}
func deleteCandidate(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/candidate/delete/")
	fmt.Fprintf(w, fmt.Sprintf("%q", name))
}
func voteCandidate(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	name := strings.TrimPrefix(r.URL.Path, "/candidate/vote/")
	log.Println(name)
	for _, val := range candidates {
		if val.name == name {
			val.votes += 1
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(w, fmt.Sprintf("You voted for %q", name))
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, fmt.Sprintf("%q is not a valid candidate", name))
}
func votingStatus(w http.ResponseWriter, r *http.Request) {
	var sb strings.Builder
	for _, candidate := range candidates {
		sb.WriteString(candidate.name + " ")
		sb.WriteString(strconv.Itoa(candidate.votes) + "\n")
	}
	fmt.Fprintf(w, sb.String())
}
func votingResult(w http.ResponseWriter, r *http.Request) {
	var max int
	var winner string
	var check bool
	for _, candidate := range candidates {
		if candidate.votes > max {
			winner = candidate.name
			max = candidate.votes
		} else if candidate.votes == max && max != 0 {
			check = true
			winner = winner + " " + candidate.name
		}
	}
	if winner == "" {
		w.WriteHeader(http.StatusExpectationFailed)
		fmt.Fprintf(w, fmt.Sprintf("No votes are casted. No candidate is declared as winner"))
	} else {
		if check {
			w.WriteHeader(http.StatusConflict)
			winner = strings.Join(strings.Split(winner, " "), ",")
			fmt.Fprintf(w, fmt.Sprintf("Draw between %q with %d votes", winner, max))
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, fmt.Sprintf("The winner of the elections is %q with %d votes", winner, max))
		}
	}
}
