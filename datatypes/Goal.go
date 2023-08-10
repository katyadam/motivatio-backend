package datatypes

import "time"

// Goal
/**
	This datatype should match db table as well
**/
type Goal struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	Relevancy   float32   `json:"relevance"`
	Pin         bool      `json:"pin"`
	Done        bool      `json:"done"`
}
