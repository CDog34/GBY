package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	. "github.com/CDog34/GBY/server/models"
	"time"
	"log"
	"gopkg.in/mgo.v2/bson"
)

func ArticleIndex(w http.ResponseWriter, r *http.Request) {
	test := Article{}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(test.List()); err != nil {
		panic(err)
	}
}

func ArticleShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	article := Article{}
	encoder := json.NewEncoder(w)
	err := article.GetOne(articleId)
	if err != nil {
		log.Print(err)
		if err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)
			encoder.Encode(map[string]interface{}{"status":false, "message":"not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(map[string]interface{}{"status":false})
		}
	} else {
		w.WriteHeader(http.StatusOK)
		encoder.Encode(article)
	}
}

func ArticleCreate(w http.ResponseWriter, r *http.Request) {
	newArticle := Article{
		Title:"喵喵喵",
		UpdateAt:time.Now(),
		Author:"CDog",
		Content:"fdsfasdfaswerfohiszhfjl;askdjfoiasuerfopuaehrfo;vilqeuy4rfoisargheudpifoiaerhslifo;iawquheblrjfh",
	}
	err := newArticle.Save()
	encoder := json.NewEncoder(w);
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(map[string]interface{}{"status":false})
	} else {
		w.WriteHeader(http.StatusOK)
		encoder.Encode(newArticle)
	}
}

func ArticleUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	newArticle := Article{
		UpdateAt:time.Now(),
	}
	newArticle.Id = bson.ObjectIdHex(articleId)
	err := newArticle.Save()
	encoder := json.NewEncoder(w)
	if err != nil {
		log.Print(err)
		if err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)
			encoder.Encode(map[string]interface{}{"status":false, "message":"not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(map[string]interface{}{"status":false})
		}
	} else {
		w.WriteHeader(http.StatusOK)
		encoder.Encode(newArticle)
	}
}

func ArticleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	article := Article{}
	err := article.MarkDeleted(articleId)
	encoder := json.NewEncoder(w)
	if err != nil {
		log.Print(err)
		if err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)
			encoder.Encode(map[string]interface{}{"status":false, "message":"not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(map[string]interface{}{"status":false})
		}
	} else {
		w.WriteHeader(http.StatusOK)
		encoder.Encode(article)
	}
}