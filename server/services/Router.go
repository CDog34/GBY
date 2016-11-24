package services

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
	"github.com/CDog34/GBY/server/errors"
)

type JsonHandlerFunc func(http.ResponseWriter, *http.Request) (error, interface{})

type Route struct {
	Name           string
	Method         string
	Pattern        string
	HandleFunc     JsonHandlerFunc
	PureHandleFunc http.HandlerFunc
	Queries        []string
	NoJson         bool
	handler        http.HandlerFunc
}
type Routes []Route

func completeRequest(err error, okData interface{}, w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	if err != nil {
		log.Print(err)
		errorDetail := errorHandler.GetErrorDetail(err)
		w.WriteHeader(errorDetail.HttpCode)
		encoder.Encode(errorDetail)
	} else {
		if okData == nil {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusOK)
			encoder.Encode(okData)
		}

	}
}

func wrapJsonHandler(inner JsonHandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		err, rst := inner(w, r)
		completeRequest(err, rst, w, r)
	})
}

func NewRouter(appRoute *Routes) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range *appRoute {

		if !route.NoJson {
			route.handler = wrapJsonHandler(route.HandleFunc)
		} else {
			route.handler = http.HandlerFunc(route.PureHandleFunc)
		}

		rule := router.
		Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(Logger(route.handler, route.Name))
		if route.Queries != nil {
			rule.Queries(route.Queries...)
		}
		log.Printf("Setup Route: %s [%s]", route.Name, route.Pattern)
	}
	return router
}