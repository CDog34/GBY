package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	. "github.com/CDog34/GBY/server/models"
)

func ArticleIndex(w http.ResponseWriter, r *http.Request) {
	todos := Articles{
		Article{Title:"TestArticle"},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func ArticleShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	json.NewEncoder(w).Encode(map[string]interface{}{"status":true, "articleId":articleId})
}