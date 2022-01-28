package handlers

import (
	"fmt"
	"net/http"

	"github.com/osamamosaad/url-shortener-api-ugnxea/pkg/models"
)

type Encode struct {
	Db models.UrlInterface
}

func (e *Encode) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}
