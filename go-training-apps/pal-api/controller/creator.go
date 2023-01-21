package controller

import (
	"encoding/json"
	"net/http"
	"pal-api/views"
)

func creator() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		data := views.Response{
			Code: http.StatusOK,
			Body: "Paddy `O LaG is my creator!!! ;)",
		}
		json.NewEncoder(w).Encode(data)

	}
}
