package Model

import (
	"errors"
	. "github.com/CDog34/GBY/server/services"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const linkCollectionName = "link"

type Link struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name"`
	Url         string        `json:"Url"`
	UpdateAt    time.Time     `json:"updateAt"`
	IsPublic    bool          `json:"isPublic"`
	Image       string        `json:"image"`
	Description string        `json:"description "`
}

type Links []Link

func (a *Link) List(all bool) (error, Links) {
	db := &DBService
	var query *mgo.Query
	var dbSession *DBSession
	if all {
		query, dbSession = db.Retrieve(linkCollectionName, nil)
	} else {
		query, dbSession = db.Retrieve(linkCollectionName, bson.M{"isPublic": true})
	}
	defer dbSession.Close()
	result := make(Links, 0, 10)
	err := query.Sort("-_id").All(&result)
	return err, result
}

func (a *Link) Save() error {
	db := &DBService
	if a.Id.Valid() {
		res, dbs := db.Update(linkCollectionName, bson.M{"_id": a.Id}, a)
		defer dbs.Close()
		return res
	} else {
		a.Id = bson.NewObjectId()
		res, dbs := db.Create(linkCollectionName, a)
		defer dbs.Close()
		return res
	}
}

func (a *Link) GetOne(linkId string) error {
	db := &DBService
	if !bson.IsObjectIdHex(linkId) {
		return errors.New("paramErr.inValidObjectId/*/" + linkId)
	}
	res, dbs := db.Retrieve(linkCollectionName, bson.M{"_id": bson.ObjectIdHex(linkId)})
	defer dbs.Close()
	return res.One(a)
}

func (a *Link) MarkDeleted(linkId string) error {
	err := a.GetOne(linkId)
	if err != nil {
		return err
	}
	a.IsPublic = false
	return a.Save()
}

func (a *Link) RecoverDeleted(linkId string) error {
	err := a.GetOne(linkId)
	if err != nil {
		return err
	}
	a.IsPublic = true
	return a.Save()
}

func (a *Link) HardDelete(linkId string) error {
	db := &DBService
	if !bson.IsObjectIdHex(linkId) {
		return errors.New("paramErr.inValidObjectId/*/" + linkId)
	}
	res, dbs := db.Delete(linkCollectionName, bson.M{"_id": bson.ObjectIdHex(linkId)})
	defer dbs.Close()
	return res
}
