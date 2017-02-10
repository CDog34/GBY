package handlers

import (
	"errors"
	. "github.com/CDog34/GBY/server/models"
	"github.com/CDog34/GBY/server/services"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"time"
)

func ArticleIndex(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
		return err, nil
	}
	onIndex := len(r.Form["onIndex"]) > 0 && r.Form["onIndex"][0] == "1"
	article := Article{}
	return article.List(false, onIndex)
}
func ArticleListAll(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	article := Article{}
	return article.List(true, true)
}

var articlePostRules = services.FieldRules{
	"title": services.FieldRule{
		"required": true,
		"type":     "string",
	},
	"author": services.FieldRule{
		"required": true,
		"type":     "string",
	},
	"content": services.FieldRule{
		"required": true,
		"type":     "string",
	},
	"showOnIndex": services.FieldRule{
		"required": false,
		"type":     "boolean",
	},
	"updateAt": services.FieldRule{
		"required": false,
		"type":     "string",
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
	params := services.PostParams{Request: r, Rules: articlePostRules}
	if err := params.Valid(); err != nil {
		return err, nil
	}
	newTime, parseErr := time.Parse("2006-01-02T15:04:05.000Z", params.GetString("updateAt"))
	if parseErr != nil {
		newTime = time.Now()
	}
	newArticle := Article{
		Title:    params.GetString("title"),
		UpdateAt: newTime,
		Author:   params.GetString("author"),
		Content:  params.GetString("content"),
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
	params := services.PostParams{Request: r, Rules: articlePostRules}
	if err := params.Valid(); err != nil {
		return err, nil
	}
	newTime, parseErr := time.Parse("2006-01-02T15:04:05.000Z", params.GetString("updateAt"))
	if parseErr != nil {
		newTime = time.Now()
	}
	newArticle := Article{
		Title:       params.GetString("title"),
		UpdateAt:    newTime,
		Author:      params.GetString("author"),
		Content:     params.GetString("content"),
		ShowOnIndex: params.GetBool("showOnIndex"),
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
func ArticleRecover(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	article := Article{}
	err := article.RecoverDeleted(articleId)
	return err, nil
}
