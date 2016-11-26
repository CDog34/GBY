package main

import (
	"github.com/CDog34/GBY/server/configs"
	. "github.com/CDog34/GBY/server/routes"
	. "github.com/CDog34/GBY/server/services"
	"log"
	"net/http"
)

func main() {
	log.Println("Server start!")
	router := NewRouter(&RouteRules)
	log.Printf("Server start on %s", config.Get("listenAddr").(string))
	log.Fatal(http.ListenAndServe(config.Get("listenAddr").(string), router))
}
