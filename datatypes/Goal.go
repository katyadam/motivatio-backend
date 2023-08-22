package datatypes

// Goal
/**
	This datatype should match db table as well
**/
type Goal struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	StartDate   string  `json:"startDate"`
	Relevancy   float32 `json:"relevancy"`
	Pin         bool    `json:"pin"`
	Done        bool    `json:"done"`
	UserId      string  `json:"userId"`
}
