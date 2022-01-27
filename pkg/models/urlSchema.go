package models

type Url struct {
	ID      int    `json:"id"`
	Short   string `json:"shortKey"`
	Orignal string `json:"orignalUrl"`
}
