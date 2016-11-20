package services

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandleFunc  http.HandlerFunc
	Queries     []string
	ContentType string
}
type Routes []Route

func NewRouter(appRoute *Routes) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range *appRoute {

		handler := route.HandleFunc
		handlerFunc := Logger(handler, route.Name)
		handlerFuncProceed := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if len(route.ContentType) != 0 {
				w.Header().Set("Content-Type", route.ContentType)
			} else {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			}
			handlerFunc.ServeHTTP(w, r)
		})

		rule := router.
		Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handlerFuncProceed)
		if route.Queries != nil {
			rule.Queries(route.Queries...)
		}
	}
	return router
}