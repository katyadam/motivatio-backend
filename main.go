package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Goal struct {
	ID          int     `json:"id"`
	Title       string  `json:"name"`
	Description string  `json:"description"`
	Relevance   float32 `json:"relevance"`
}

func goalsHandler(c *gin.Context) {
	var goals []Goal
	goals = append(goals, Goal{ID: 1, Title: "Gym", Description: "descripce", Relevance: 3.5})
	goals = append(goals, Goal{ID: 2, Title: "Uceni", Description: "descripce", Relevance: 4.5})
	goals = append(goals, Goal{ID: 3, Title: "Beh", Description: "descripce", Relevance: 1.5})
	goals = append(goals, Goal{ID: 4, Title: "Socializace", Description: "descripce", Relevance: 8.5})

	c.JSON(200, gin.H{
		"goals": goals,
	})
}

func main() {

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/goals", goalsHandler)
	r.Run(":5000")
}
