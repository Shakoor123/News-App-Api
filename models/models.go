package models

type User struct {
	Uid int `json:"uid,omitempty"`
	Username string `json:"username,omitempty"`
	Password string	`json:"password,omitempty"`
	Email string	`json:"email,omitempty"`

}

type News struct {
	Nid int `json:"nid,omitempty"`
	Uid int `json:"uid,omitempty"`
	Title string `json:"title,omitempty"`
	Place string	`json:"place,omitempty"`
	Description string	`json:"description,omitempty"`
	Date string	`json:"date,omitempty"`

}