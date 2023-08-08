package datatypes

import (
	"time"
)

// Goal
/**
	This datatype should match db table as well
**/
type Goal struct {
	ID          int       `json:"id"`
	Title       string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"date"`
	Relevancy   float32   `json:"relevance"`
}
