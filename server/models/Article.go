package Model

import (
	"time"
	. "github.com/CDog34/GBY/server/services"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

const articleCollectionName = "article"

type Article struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title    string `json:"title"`
	UpdateAt time.Time `json:"updateAt"`
	Author   string `json:"author"`
	Content  string `json:"content"`
	Deleted  bool `json:"deleted"`
}

type Articles []Article

func (a *Article) List() (error, Articles) {
	db := &DBService
	query := db.Retrieve(articleCollectionName, nil)
	//defer db.Close()
	result := make(Articles, 0, 10)
	err := query.All(&result)
	return err, result
}

func (a *Article) Save() error {
	db := &DBService
	if a.Id.Valid() {
		return db.Update(articleCollectionName, bson.M{"_id":a.Id}, a)
	} else {
		a.Id = bson.NewObjectId()
		return db.Create(articleCollectionName, a)
	}
}

func (a *Article) GetOne(articleId string) error {
	db := &DBService
	if !bson.IsObjectIdHex(articleId) {
		return errors.New("Invalid ObjectId")
	}
	return db.Retrieve(articleCollectionName, bson.M{"_id":bson.ObjectIdHex(articleId)}).One(a)
}

func (a *Article) MarkDeleted(articleId string) error {
	err := a.GetOne(articleId)
	if err != nil {
		return err
	}
	a.Deleted = true
	return a.Save()
}

func (a *Article) RecoverDeleted(articleId string) error {
	err := a.GetOne(articleId)
	if err != nil {
		return err
	}
	a.Deleted = false
	return a.Save()
}

func (a *Article) HardDelete(articleId string) error {
	db := &DBService
	if !bson.IsObjectIdHex(articleId) {
		return errors.New("Invalid ObjectId")
	}
	return db.Delete(articleCollectionName, bson.M{"_id":bson.ObjectIdHex(articleId)})
}
