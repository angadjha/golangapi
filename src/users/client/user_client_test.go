package client

import (
	"encoding/json"
	"golangapi/src/users/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	// Create a mock server
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reqBody models.UserRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := models.UserResponse{
			Status: "20000",
		}

		respJSON, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	apiURL := ts.URL
	authToken := "testToken"

	userInfo := models.UserRequest{
		Name:   "Angad",
		Salary: "23000",
		Age:    "25",
	}

	resp, err := CreateUser(apiURL, authToken, userInfo)
	if err != nil {
		t.Fatalf("Error checking eligibility: %v", err)
	}

	if resp.User.Name != "Jha" {
		t.Error("Name is not match")
	}

	if resp.User.Salary != "30000" {
		t.Errorf("Expected Salary: %s", resp.User.Salary)
	}
}
