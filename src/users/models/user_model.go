package models

type UserResponse struct {
	Status string `json:"status"`
	User   User   `json:"data"`
}

type User struct {
	Name   string `json:"name"`
	Salary string `json:"salary"`
	Age    string `json:"age"`
}

type UserRequest struct {
	Name   string `json:"name"`
	Salary string `json:"salary"`
	Age    string `json:"age"`
}
