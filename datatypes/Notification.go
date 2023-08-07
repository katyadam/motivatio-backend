package datatypes

type TypeEnum string

const (
	Reminder  TypeEnum = "Reminder to take action!"
	Milestone TypeEnum = "You reach new milestone!"
)

type Notification struct {
	ID   int      `json:"id"`
	Type TypeEnum `json:"type"`
}
