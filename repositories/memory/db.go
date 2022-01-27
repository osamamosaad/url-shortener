package memory

import "github.com/osamamosaad/url-shortener-api-ugnxea/pkg/models"

var db = make(map[string]*models.Url)
