package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	. "github.com/CDog34/GBY/server/models"
	"time"
	"log"
	"gopkg.in/mgo.v2/bson"
	"github.com/CDog34/GBY/server/services"
	"errors"
)

func ArticleIndex(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	article := Article{}
	return article.List()
}

var articlePostRules = services.FieldRules{
	"title":services.FieldRule{
		"required":true,
		"type":"string",
	},
	"author":services.FieldRule{
		"required":true,
		"type":"string",
	},
	"content":services.FieldRule{
		"required":true,
		"type":"string",
	},
}

func ArticleShow(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	article := Article{}
	err := article.GetOne(articleId)
	return err, article
}

func ArticleCreate(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	params := services.PostParams{Request:r, Rules: articlePostRules}
	if err := params.Valid(); err != nil {
		return err, nil
	}
	newArticle := Article{
		Title:params.GetString("title"),
		UpdateAt:time.Now(),
		Author:params.GetString("author"),
		Content:params.GetString("content"),
	}
	err := newArticle.Save()
	return err, newArticle
}

func ArticleUpdate(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	if !bson.IsObjectIdHex(articleId) {
		return errors.New("paramErr.inValidObjectId/*/" + articleId), nil
	}
	params := services.PostParams{Request:r, Rules: articlePostRules}
	if err := params.Valid(); err != nil {
		return err, nil
	}
	newArticle := Article{
		Title:params.GetString("title"),
		UpdateAt:time.Now(),
		Author:params.GetString("author"),
		Content:params.GetString("content"),
	}
	newArticle.Id = bson.ObjectIdHex(articleId)
	err := newArticle.Save()
	return err, newArticle
}

func ArticleDelete(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	vars := mux.Vars(r)
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}
	isHard := len(r.Form["hard"]) > 0 && r.Form["hard"][0] == "1"
	articleId := vars["articleId"]
	article := Article{}
	if !isHard {
		err = article.MarkDeleted(articleId)
	} else {
		err = article.HardDelete(articleId)
	}
	return err, nil
}