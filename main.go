package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Goal struct {
	ID          int     `json:"id"`
	Title       string  `json:"name"`
	Description string  `json:"description"`
	Relevance   float32 `json:"relevance"`
}

var goals []Goal

func getItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goals)
}

func main() {
	goals = append(goals, Goal{ID: 1, Title: "Gym", Description: "descripce", Relevance: 3.5})
	goals = append(goals, Goal{ID: 2, Title: "Uceni", Description: "descripce", Relevance: 4.5})
	goals = append(goals, Goal{ID: 3, Title: "Beh", Description: "descripce", Relevance: 1.5})
	goals = append(goals, Goal{ID: 4, Title: "Socializace", Description: "descripce", Relevance: 8.5})

	http.HandleFunc("/api/goals", getItemsHandler)

	fmt.Println("Listen for api on http://localhost:8080/api/goals")
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Server shutting down...")
}
