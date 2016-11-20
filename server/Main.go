package main

import (
	"log"
	"net/http"
	. "github.com/CDog34/GBY/server/services"
)

func main() {
	router := NewRouter(&routes)
	log.Fatal(http.ListenAndServe(":8080", router))
}
