package Route

import (
	"encoding/json"
	"github.com/CDog34/GBY/server/errors"
	"github.com/CDog34/GBY/server/models"
	. "github.com/CDog34/GBY/server/services"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"net/http"
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
	Roles          []int
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
func checkRoleWrapper(handlerFunc JsonHandlerFunc, roles []int) JsonHandlerFunc {
	if roles == nil {
		return handlerFunc
	}
	return func(w http.ResponseWriter, r *http.Request) (error, interface{}) {
		sess, err := SessionMgr.SessionStart(w, r, false)
		if err != nil {
			return err, nil
		}
		if sess.Get("user") == nil {
			return errors.New("auth.notLogin"), nil
		}
		allow := false
		for _, val := range roles {
			if val == sess.Get("user").(Model.User).Role {
				allow = true
				break
			}
		}
		if !allow {
			return errors.New("auth.notAllow"), nil
		}
		servErr, servRst := handlerFunc(w, r)
		return servErr, servRst
	}
}

func NewRouter(appRoute *Routes) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range *appRoute {

		if !route.NoJson {
			route.handler = wrapJsonHandler(checkRoleWrapper(route.HandleFunc, route.Roles))
		} else {
			route.handler = http.HandlerFunc(route.PureHandleFunc)
		}

		rule := router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(Logger(CrossDomainHeaders(route.handler), route.Name))
		if route.Queries != nil {
			rule.Queries(route.Queries...)
		}
		router.
			Methods("OPTIONS").
			Path(route.Pattern).
			Name(route.Name + "opt").
			Handler(CrossDomainHeaders(func(res http.ResponseWriter, req *http.Request) {
				res.WriteHeader(http.StatusOK)
				return
			}))
		log.Printf("Setup Route: %s [%s]", route.Name, route.Pattern)
	}
	return router
}
