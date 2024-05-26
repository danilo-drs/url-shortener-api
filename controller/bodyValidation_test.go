package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"meli-api/model"
	"net/http"
	"testing"
)

func TestSuccessValidateRequestBody(t *testing.T) {
	// Create a sample request body
	requestBody := model.ShortUrl{
		OriginalURL: "https://example.com",
	}

	// Convert the request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Create a new HTTP request with the JSON body
	req, err := http.NewRequest("POST", "/api/shorten", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}

	// Call the validateRequestBody function
	statusCode, shortUrl, err := validateRequestBody(req)

	// Check the status code
	if statusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, statusCode)
	}

	// Check the shortUrl object
	if shortUrl.OriginalURL != requestBody.OriginalURL {
		t.Errorf("Expected original URL %s, but got %s", requestBody.OriginalURL, shortUrl.OriginalURL)
	}

	// Check the error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

type errorReader struct{}

func (er *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("error reading body")
}

func TestErrorReadingBody(t *testing.T) {
	// Create a new HTTP request with an invalid body

	req, _ := http.NewRequest("POST", "/api/shorten", &errorReader{})

	// Call the validateRequestBody function
	statusCode, _, err := validateRequestBody(req)

	// Check the status code
	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, statusCode)
	}

	// Check the error
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

func TestErrorParsingBody(t *testing.T) {
	// Create a new HTTP request with an invalid body
	req, _ := http.NewRequest("POST", "/api/shorten", bytes.NewBufferString(""))

	// Call the validateRequestBody function
	statusCode, _, err := validateRequestBody(req)

	// Check the status code
	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, statusCode)
	}

	// Check the error
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

func TestMissingOriginalURL(t *testing.T) {
	// Create a sample request body with missing OriginalURL
	requestBody := model.ShortUrl{}

	// Convert the request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Create a new HTTP request with the JSON body
	req, err := http.NewRequest("POST", "/api/shorten", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}

	// Call the validateRequestBody function
	statusCode, _, err := validateRequestBody(req)

	// Check the status code
	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, statusCode)
	}

	// Check the error
	if err == nil {
		t.Errorf("Expected error, but got %v", err)
	}
}
