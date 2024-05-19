package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"meli-api/model"
	"net/http"
)

func validateRequestBody(r *http.Request) (statusCode int, shortUrl model.ShortUrl, err error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading request body: %v", err)
		return http.StatusInternalServerError, shortUrl, err
	}
	// parse the request body
	err = json.Unmarshal(body, &shortUrl)
	if err != nil {
		fmt.Printf("Error parsing request body: %v", err)
		return http.StatusInternalServerError, shortUrl, err
	}

	// validate the original URL exists
	if shortUrl.OriginalURL == "" {
		fmt.Print("Original URL is required")
		return http.StatusInternalServerError, shortUrl, err
	}

	// return 200 if the request is valid
	return http.StatusOK, shortUrl, nil
}
