package memory

import (
	"errors"

	"github.com/osamamosaad/url-shortener-api-ugnxea/pkg/models"
)

type UrlRepo struct {
}

func (u *UrlRepo) Insert(url models.Url) (models.Url, error) {
	db[url.Short] = &url
	return u.Get(url.Short)
}

func (u UrlRepo) Get(shortKey string) (models.Url, error) {
	url, ok := db[shortKey]
	if !ok {
		return models.Url{}, errors.New("not found")
	}

	return *url, nil
}
