package controller

import (
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// validate the request body
	statusCode, shortUrl, err := validateRequestBody(r)
	if err != nil {
		w.WriteHeader(statusCode)
		w.Write([]byte(`{"success": false, "message": "Error validating request: ` + err.Error() + `"}`))
		return
	}

	// generate short URL
	shortUrl.ShortUrl, err = shortUrl.GenerateShortUrl()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"success": false, "message": "Error generating short URL: ` + err.Error() + `"}`))
		return
	}

	// Return the short URL
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": true, "shortUrl": "` + shortUrl.ShortUrl + `"}`))
}
