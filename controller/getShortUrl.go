package controller

import (
	"net/http"

	"meli-api/model"

	"github.com/gorilla/mux"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// get the key from the URL
	key := mux.Vars(r)["key"]
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"success": false, "message": "Key is required"}`))
		return
	}

	// get the short URL
	shortUrl := model.ShortUrl{}
	found, err := shortUrl.FillFromKey(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"success": false, "message": "Error getting short URL: ` + err.Error() + `"}`))
		return
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"success": false, "message": "Short URL not found"}`))
		return
	}

	// return the short URL

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": true, "shortUrl": ` + shortUrl.String() + `}`))
}
