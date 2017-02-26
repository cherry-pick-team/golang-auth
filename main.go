package main

import (
	"log"
	"net/http"
	"golang-auth/src/Router"
)

func main() {

	router := Router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
