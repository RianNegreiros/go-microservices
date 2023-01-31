package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "Broker",
	}

	_ = app.writeJSON(w, r, http.StatusOK, payload)
}
