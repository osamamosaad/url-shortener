package handlers

import (
	"encoding/json"
	"net/http"
)

type request struct {
	Url string `json:"url"`
}

type errorResponse struct {
	Error string `json:"erro"`
}

func responseWriter(status int, data interface{}, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(data)
	w.Write(jsonResp)
}
