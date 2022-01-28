package main

import (
	"log"
	"net/http"

	"github.com/osamamosaad/url-shortener-api-ugnxea/config"
	"github.com/osamamosaad/url-shortener-api-ugnxea/handlers"
	"github.com/osamamosaad/url-shortener-api-ugnxea/repositories/memory"
)

func main() {
	log.Println("start http server")
	muxServ := http.NewServeMux()
	muxServ.Handle("/encode", &handlers.Encode{&memory.UrlRepo{}})
	muxServ.Handle("/decode/", &handlers.Decode{&memory.UrlRepo{}})

	log.Fatal(http.ListenAndServe(config.PORT, muxServ))
}
