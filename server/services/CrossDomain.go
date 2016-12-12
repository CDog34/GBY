package services

import (
	"github.com/CDog34/GBY/server/configs"
	"net/http"
)

func CrossDomainHeaders(inner http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-App-Id")
		w.Header().Set("Access-Control-Allow-Origin", config.Get("accessControlAllowOriginHeader").(string))
		inner.ServeHTTP(w, r)
	})
}
