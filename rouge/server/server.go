package main

import (
	"log"
	"net/http"
)

const (
	Addsrv       = ":4040"
	TemplatesDir = "."
)

func main() {
	fileSrv := http.FileServer(http.Dir(TemplatesDir))

	if err := http.ListenAndServe(Addsrv, fileSrv); err != nil {
		log.Fatalln(err)
	}
}
