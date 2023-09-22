package client

import (
	"encoding/json"
	"fmt"
	"golangapi/src/users/models"
	"net/http"
)

func CreateUser(apiURL, authToken string, userInfo models.UserRequest) (*models.UserResponse, error) {
	req, err := CreatePostRequest(apiURL, authToken, userInfo)
	if err != nil {
		return nil, err
	}

	resp, err := SendRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var userResponse models.UserResponse
	err = json.NewDecoder(resp.Body).Decode(&userResponse)
	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}
