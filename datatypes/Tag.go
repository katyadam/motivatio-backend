package datatypes

//Tag
/**
	This datatype should match db table as well
**/
type Tag struct {
	ID      int    `json:"id"`
	TagName string `json:"name"`
	Color   int    `json:"color"`
}
