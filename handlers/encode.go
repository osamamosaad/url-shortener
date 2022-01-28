package handlers

import (
	"encoding/json"
	"net/http"

	generator "github.com/osamamosaad/url-shortener-api-ugnxea/pkg/Generator"
	"github.com/osamamosaad/url-shortener-api-ugnxea/pkg/counter"
	"github.com/osamamosaad/url-shortener-api-ugnxea/pkg/models"
)

type Encode struct {
	Db models.UrlInterface
}

func (e *Encode) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request request
	bodyIo := json.NewDecoder(r.Body)
	if err := bodyIo.Decode(&request); err != nil {
		responseWriter(http.StatusBadRequest, errorResponse{"Bad request!"}, w)
		return
	}

	url, err := e.insert(request)

	if err != nil {
		responseWriter(
			http.StatusInternalServerError,
			errorResponse{"Something went wrong!, please try after while."},
			w,
		)
		return
	}
	urlResponseBuilder(url, w)
}

func (e *Encode) insert(request request) (models.Url, error) {
	// Get the next number from counter service
	id := counter.Get()
	// Generate short key
	shortCode := generator.Create(id)

	// Build model
	url := models.Url{
		ID:      id,
		Short:   shortCode,
		Orignal: request.Url,
	}
	return e.Db.Insert(url)
}
