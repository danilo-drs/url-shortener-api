package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// validate the request body
	statusCode, shortUrl, err := validateRequestBody(r)
	if err != nil {
		w.WriteHeader(statusCode)
		w.Write([]byte(`{"success": false, "message": "Error validating request: ` + err.Error() + `"}`))
		return
	}
	inputUrl := shortUrl.OriginalURL

	// get the key from the URL
	key := mux.Vars(r)["key"]
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"success": false, "message": "Key is required"}`))
		return
	}

	// fill the short URL from the key
	found, err := shortUrl.FillFromKey(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"success": false, "message": "Error getting short URL: ` + err.Error() + `"}`))
		return
	}
	// return 404 if the short URL is not found
	if !found {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"success": false, "message": "Short URL not found"}`))
		return
	}

	// update the short URL
	shortUrl.OriginalURL = inputUrl
	err = shortUrl.Update()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"success": false, "message": "Error updating short URL: ` + err.Error() + `"}`))
		return
	}

	// return the short URL
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": true, "shortUrl": "` + shortUrl.ShortUrl + `"}`))
}
