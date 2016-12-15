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

func LinkList(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	link := Link{}
	return link.List(false)
}
func LinkListAll(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	link := Link{}
	return link.List(true)
}

var linkPostRules = services.FieldRules{
	"name": services.FieldRule{
		"required": true,
		"type":     "string",
	},
	"url": services.FieldRule{
		"required": true,
		"type":     "string",
	},
	"image": services.FieldRule{
		"required": true,
		"type":     "string",
	},
	"description": services.FieldRule{
		"required": true,
		"type":     "string",
	},
}

func LinkShow(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	vars := mux.Vars(r)
	linkId := vars["linkId"]
	link := Link{}
	err := link.GetOne(linkId)
	return err, link
}

func LinkCreate(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	params := services.PostParams{Request: r, Rules: linkPostRules}
	if err := params.Valid(); err != nil {
		return err, nil
	}
	newLink := Link{
		Name:        params.GetString("name"),
		UpdateAt:    time.Now(),
		Url:         params.GetString("url"),
		Image:       params.GetString("image"),
		Description: params.GetString("description"),
	}
	err := newLink.Save()
	return err, newLink
}

func LinkUpdate(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	vars := mux.Vars(r)
	LinkId := vars["linkId"]
	if !bson.IsObjectIdHex(LinkId) {
		return errors.New("paramErr.inValidObjectId/*/" + LinkId), nil
	}
	params := services.PostParams{Request: r, Rules: linkPostRules}
	if err := params.Valid(); err != nil {
		return err, nil
	}
	newLink := Link{
		Name:        params.GetString("name"),
		UpdateAt:    time.Now(),
		Url:         params.GetString("url"),
		Image:       params.GetString("image"),
		Description: params.GetString("description"),
	}
	newLink.Id = bson.ObjectIdHex(LinkId)
	err := newLink.Save()
	return err, newLink
}

func LinkDelete(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	vars := mux.Vars(r)
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}
	isHard := len(r.Form["hard"]) > 0 && r.Form["hard"][0] == "1"
	linkId := vars["linkId"]
	link := Link{}
	if !isHard {
		err = link.MarkDeleted(linkId)
	} else {
		err = link.HardDelete(linkId)
	}
	return err, nil
}
func LinkRecover(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	vars := mux.Vars(r)
	linkId := vars["linkId"]
	link := Link{}
	err := link.RecoverDeleted(linkId)
	return err, nil
}
