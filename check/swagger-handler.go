package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./check/swaggerui"))
	http.Handle("/v1/swagger/", http.StripPrefix("/v1/swagger/", fs))

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal(err)
	}
}