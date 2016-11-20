package main

import (
	"log"
	"net/http"
	. "github.com/CDog34/GBY/server/services"
	. "github.com/CDog34/GBY/server/routes"
)

func main() {
	router := NewRouter(&RouteRules)
	log.Fatal(http.ListenAndServe(":8080", router))
}
