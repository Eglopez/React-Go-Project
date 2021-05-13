package models

type Tag struct {
	Id    int    `json:"id"`
	Name  string `json:"tagName"`
	Color string `json:"tagColor"`
}
