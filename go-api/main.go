package main

import (
	"net/http"
	"sync"
)

type Score struct {
	Name string `json:"name"`
	Score int `json:"score"`
	Time string `json:"time"`
}

var (
	scores []Score
	mutex sync.Mutex
)

func main () {
	// http.HandleFunc("/score", handleScore)
	// http.HandleFunc("/scores", getScores) 
}

func handleScore(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var newScore Score 
	if err := j
}