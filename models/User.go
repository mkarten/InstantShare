package models

type User struct {
	UUID      string `json:"uuid"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
