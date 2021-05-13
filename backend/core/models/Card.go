package models

type Card struct {
	Id          int    `json:"id"`
	Name        string `json:"cardName"`
	Description string `json:"cardDescription"`
	Status      string `json:"status"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}
