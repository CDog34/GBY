package handlers

import (
	"net/http"
	"encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {

	if err := json.NewEncoder(w).Encode(map[string]string{"system": "ok"}); err != nil {
		panic(err)
	}
}