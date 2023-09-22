package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePostRequest(t *testing.T) {
	apiURL := "https://dummy.restapiexample.com/api/v1/create"
	authToken := "testToken"
	body := map[string]interface{}{"key": "value"}

	req, err := CreatePostRequest(apiURL, authToken, body)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	if req.Method != http.MethodPost {
		t.Errorf("Expected POST method, got %s", req.Method)
	}

	if req.Header.Get("Authorization") != "Bearer "+authToken {
		t.Errorf("Authorization header is incorrect")
	}
}

func TestSendRequest(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	resp, err := SendRequest(req)
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
