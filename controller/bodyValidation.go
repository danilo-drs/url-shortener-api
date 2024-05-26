package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"meli-api/model"
	"net/http"
)

func validateRequestBody(r *http.Request) (statusCode int, shortUrl model.ShortUrl, err error) {

	body, er := io.ReadAll(r.Body)
	if er != nil {
		fmt.Printf("Error reading request body: %v", er)
		return http.StatusBadRequest, shortUrl, er
	}
	// parse the request body
	er = json.Unmarshal(body, &shortUrl)
	if er != nil {
		fmt.Printf("Error parsing request body: %v", er)
		return http.StatusBadRequest, shortUrl, er
	}

	// validate the original URL exists
	if shortUrl.OriginalURL == "" {
		fmt.Print("Original URL is required")
		er = errors.New("original URL is required")
		return http.StatusBadRequest, shortUrl, er
	}

	// return 200 if the request is valid
	return http.StatusOK, shortUrl, nil
}
