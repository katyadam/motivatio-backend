package datatypes

type TypeEnum string

const (
	Reminder  TypeEnum = "Reminder to take action!"
	Milestone TypeEnum = "You reached new milestone!"
)

// Notification
/**
	This datatype should match db table as well
**/
type Notification struct {
	ID               int      `json:"id"`
	NotificationType TypeEnum `json:"type"`
}
