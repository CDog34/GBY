package handlers

import (
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	return nil, map[string]interface{}{
		"system": "ok",
		"time":time.Now(),
	}
}