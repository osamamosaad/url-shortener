package memory

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/osamamosaad/url-shortener-api-ugnxea/pkg/models"
)

type UrlRepo struct {
}

func (u *UrlRepo) Insert(url models.Url) (models.Url, error) {
	s := url.Short + strconv.Itoa(url.ID)
	db[s] = &url
	fmt.Println(db)
	return u.Get(s)
}

func (u UrlRepo) Get(shortKey string) (models.Url, error) {
	url, ok := db[shortKey]
	if !ok {
		return models.Url{}, errors.New("not found")
	}

	return *url, nil
}
