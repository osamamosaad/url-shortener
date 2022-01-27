package models

type UrlInterface interface {
	Insert(Url) (Url, error)
	Get(string) (Url, error)
}
