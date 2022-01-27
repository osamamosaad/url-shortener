package handlers

import (
	"fmt"
	"net/http"
)

type Encode struct {
}

func (e *Encode) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}
