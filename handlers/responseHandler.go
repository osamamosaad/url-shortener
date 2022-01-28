package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/osamamosaad/url-shortener-api-ugnxea/helpers"
	"github.com/osamamosaad/url-shortener-api-ugnxea/pkg/models"
)

type request struct {
	Url string `json:"url"`
}

type urlResponse struct {
	ID       int    `json:"id"`
	ShortUrl string `json:"shortUrl,omitempty"`
	Original string `json:"original,omitempty"`
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

func urlResponseBuilder(url models.Url, w http.ResponseWriter) {
	res := urlResponse{
		ID:       url.ID,
		ShortUrl: helpers.FullUrl(url.Short),
		Original: url.Orignal,
	}
	responseWriter(http.StatusOK, res, w)
}
