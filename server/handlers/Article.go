package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	. "github.com/CDog34/GBY/server/models"
	"time"
	"log"
	"gopkg.in/mgo.v2/bson"
	"github.com/CDog34/GBY/server/services"
)

func ArticleIndex(w http.ResponseWriter, r *http.Request) {
	article := Article{}
	err, list := article.List()
	completeRequest(err, list, w, r)
}

func ArticleShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	article := Article{}
	err := article.GetOne(articleId)
	completeRequest(err, article, w, r)
}

func ArticleCreate(w http.ResponseWriter, r *http.Request) {
	newArticle := Article{
		Title:"喵喵喵",
		UpdateAt:time.Now(),
		Author:"CDog",
		Content:"fdsfasdfaswerfohiszhfjl;askdjfoiasuerfopuaehrfo;vilqeuy4rfoisargheudpifoiaerhslifo;iawquheblrjfh",
	}
	params := services.PostParams{Request:r, Patterns: services.FieldPattern{
		"title":map[string]interface{}{
			"required":true,
		},
		"author":map[string]interface{}{
			"required":true,
		},
		"content":map[string]interface{}{
			"required":true,
		},
	}}
	if err := params.Valid(); err != nil {
		completeRequest(err, nil, w, r)
		return
	}
	err := newArticle.Save()
	completeRequest(err, newArticle, w, r)
}

func ArticleUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	newArticle := Article{
		UpdateAt:time.Now(),
	}
	newArticle.Id = bson.ObjectIdHex(articleId)
	err := newArticle.Save()
	completeRequest(err, newArticle, w, r)
}

func ArticleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}
	isHard := r.Form["hard"][0] == "1"
	articleId := vars["articleId"]
	article := Article{}
	if !isHard {
		err = article.MarkDeleted(articleId)
	} else {
		err = article.HardDelete(articleId)
	}
	completeRequest(err, nil, w, r)
}