package models

type Student struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Sex string `json:"sex"`
	Age int16 `json:"age"`
	Tel string `json:"tel"`
}
