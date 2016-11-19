package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		handler := route.HandleFunc
		handlerFunc := Logger(handler, route.Name)

		router.
		Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handlerFunc)
	}
	return router
}