package models

type Board struct {
	Id              int    `json:"id" form:"id"`
	Name            string `json:"boardName" form:"taskName"`
	BackgroundColor string `json:"backgroundColor" form:"priority"`
}
