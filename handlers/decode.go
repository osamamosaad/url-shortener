package handlers

import (
	"fmt"
	"net/http"
)

type Decode struct {
}

func (e *Decode) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}
