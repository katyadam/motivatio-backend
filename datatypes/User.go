package datatypes

import "time"

// Goal
/**
	This datatype should match db table as well
**/
type User struct {
	ID        string    `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	AddDate   time.Time `json:"addDate"`
	Phone     string    `json:"phone"`
}
