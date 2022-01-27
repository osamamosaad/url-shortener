package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("start http server")
	muxServ := http.NewServeMux()
	muxServ.Handle("/encode", &Encode{})
	muxServ.Handle("/decode", &Decode{})

	log.Fatal(http.ListenAndServe(":8000", muxServ))
}

type Encode struct {
}

func (e *Encode) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}

type Decode struct {
}

func (e *Decode) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}
