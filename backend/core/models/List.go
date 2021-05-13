package models

type List struct {
	Id     int    `json:"id"`
	Name   string `json:"listName"`
	Status string `json:"status"`
}
