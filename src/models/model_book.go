package models

type Book struct {
	ISDN string `json:"isdn"`
	Title string `json:"title"`
	Author string `json:"author"`
	Pages int `json:"page"`
}