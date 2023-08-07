package datatypes

import (
	"time"
)

type Goal struct {
	ID          int       `json:"id"`
	Title       string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Relevance   float32   `json:"relevance"`
}
