package datatypes

//Tag
/**
	This datatype should match db table as well
**/
type Tag struct {
	ID      string `json:"id"`
	TagName string `json:"name"`
	Color   int    `json:"color"`
	UserId  string `json:"userId"`
}
