package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
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
	http.HandleFunc("/score", handleScore)
	http.HandleFunc("/scores", getScores) 
	fmt.Println("Server is running on http://localhost:8080")
}

func handleScore(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var newScore Score 
	if err := json.NewDecoder(r.Body).Decode(&newScore); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	scores = append(scores, newScore)
	mutex.Unlock()
	w.WriteHeader(http.StatusCreated)
}

func getScores(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w,"Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	mutex.Lock()
	defer mutex.Unlock()

	sort.Slice(scores, func(i,j int)bool {
		return scores[i].Score > scores[j].Score
	})

	type RankedScore struct {
		Rank int `json:"rank"`
		Name string `json:"name"`
		Score int `json:"score"`
		Time string `json:"time"`
	}

	var response []RankedScore
	for i, s := range scores {
		response = append(response, RankedScore{
			Rank: i+1,
			Name: s.Name,
			Score: s.Score,
			Time: s.Time,
		})
	}
}