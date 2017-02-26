package main

import (
	"golang-auth/src/Router"
	"log"
	"net/http"
)

func main() {

	router := Router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
