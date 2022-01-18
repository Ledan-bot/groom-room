package models

type User struct {
	ID        int64  `json:"userid"`
	FirstName string `json:"fistname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
