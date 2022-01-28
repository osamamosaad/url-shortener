package handlers

import (
	"net/http"

	"github.com/osamamosaad/url-shortener-api-ugnxea/helpers"
	"github.com/osamamosaad/url-shortener-api-ugnxea/pkg/models"
)

type Decode struct {
	Db models.UrlInterface
}

func (d *Decode) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri, err := helpers.GetURI(r.URL.Path)

	if err != nil {
		responseWriter(http.StatusBadRequest, errorResponse{"short key is missing"}, w)
		return
	}

	data, err := d.Db.Get(uri)
	if err != nil {
		responseWriter(http.StatusNotFound, errorResponse{err.Error()}, w)
		return
	}
	urlResponseBuilder(data, w)
}
