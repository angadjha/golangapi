package main

import (
	"fmt"
	"golangapi/src/users/client"
	"golangapi/src/users/models"
)

func main() {
	apiURL := "https://dummy.restapiexample.com/api/v1/create"
	authToken := "YOUR_AUTH_TOKEN"

	userInfo := models.UserRequest{
		Name:   "Angad",
		Salary: "20000",
		Age:    "30",
	}

	userResponse, err := client.CreateUser(apiURL, authToken, userInfo)
	if err != nil {
		fmt.Println("Error checking eligibility:", err)
		return
	}
	fmt.Printf("Users: %v", userResponse)
	fmt.Printf("Eligible: %v, Message: %v\n", userResponse.User.Age, userResponse.User.Name)
}

type base struct {
	color string
}

func (b *base) say() {
	fmt.Printf("Hello angad")
}
