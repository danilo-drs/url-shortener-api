package controller

import (
	"net/http"

	"meli-api/model"

	"github.com/gorilla/mux"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// get the key from the URL
	key := mux.Vars(r)["key"]
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"success": false, "message": "Key is required"}`))
		return
	}

	// delete the short URL
	shortUrl := model.ShortUrl{}
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

	// delete the short URL
	err = shortUrl.Delete()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"success": false, "message": "Error deleting short URL: ` + err.Error() + `"}`))
		return
	}
}
