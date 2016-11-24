package handlers

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/CDog34/GBY/server/models"
)

func completeRequest(err error, okData interface{}, w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	if err != nil {
		log.Print(err)
		errorDetail := Model.GetErrorDetail(err)
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

func Index(w http.ResponseWriter, r *http.Request) {

	if err := json.NewEncoder(w).Encode(map[string]string{"system": "ok"}); err != nil {
		panic(err)
	}
}