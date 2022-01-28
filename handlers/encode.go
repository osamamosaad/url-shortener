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
	bodyIo := json.NewDecoder(r.Body)
	var request request
	if err := bodyIo.Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, _ := json.Marshal(errorResponse{"badRequest"})
		w.Write(jsonResp)
	}

	id := counter.Get()
	shortCode := generator.Create(id)
	url := models.Url{
		ID:      id,
		Short:   shortCode,
		Orignal: request.Url,
	}

	res, _ := e.Db.Insert(url)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(res)
	w.Write(jsonResp)
}
