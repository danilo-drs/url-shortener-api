package controller

import (
	"encoding/json"
	"net/http"

	"meli-api/model"
)

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: ADD pagination

	w.Header().Add("Content-Type", "application/json")
	// get all short URLs
	shortUrls, err := model.GetAllShortUrls()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"success": false, "message": "Error getting short URLs: ` + err.Error() + `"}`))
		return
	}

	// converte um array de ShortUrl em uma string JSON
	jsonData, err := json.Marshal(shortUrls)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"success": false, "message": "Error converting short URLs to JSON: ` + err.Error() + `"}`))
		return
	}

	// return the short URLs
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": true, "shortUrls": ` + string(jsonData) + `}`))
}
