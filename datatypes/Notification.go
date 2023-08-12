package datatypes

/*
*
Notification Enum type
*/
type NotificationType int

const (
	Reminder NotificationType = iota
	Milestone
)

type Notification struct {
	ID               int              `json:"id"`
	NotificationType NotificationType `json:"type"`
}
