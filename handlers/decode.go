package handlers

import (
	"fmt"
	"net/http"

	"github.com/osamamosaad/url-shortener-api-ugnxea/pkg/models"
)

type Decode struct {
	Db models.UrlInterface
}

func (d *Decode) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}
