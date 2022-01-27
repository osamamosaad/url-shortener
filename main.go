package main

import (
	"log"
	"net/http"

	"github.com/osamamosaad/url-shortener-api-ugnxea/handlers"
)

func main() {
	log.Println("start http server")
	muxServ := http.NewServeMux()
	muxServ.Handle("/encode", &handlers.Encode{})
	muxServ.Handle("/decode", &handlers.Decode{})

	log.Fatal(http.ListenAndServe(":8000", muxServ))
}
