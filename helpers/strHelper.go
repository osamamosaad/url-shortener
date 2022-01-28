package helpers

import (
	"errors"
	"strings"

	"github.com/osamamosaad/url-shortener-api-ugnxea/config"
)

// GetURI to get URI from path
func GetURI(path string) (string, error) {
	split := strings.Split(path, "/")
	l := len(split) - 1
	uri := strings.TrimSpace(split[l])

	if uri == "" {
		return "", errors.New("error while extracting uri")
	}

	return uri, nil
}

// FullUrl get full url
func FullUrl(urlShortKey string) string {
	return config.DOMAIN + "/" + urlShortKey
}
