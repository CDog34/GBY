package main

import (
	"log"
	"net/http"
	. "github.com/CDog34/GBY/server/services"
	. "github.com/CDog34/GBY/server/routes"
	"github.com/CDog34/GBY/server/configs"
)

func main() {
	router := NewRouter(&RouteRules)
	log.Printf("Server start on %s", config.Get("listenAddr").(string))
	log.Fatal(http.ListenAndServe(config.Get("listenAddr").(string), router))
}
