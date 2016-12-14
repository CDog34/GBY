package Model

import (
	"errors"
	. "github.com/CDog34/GBY/server/services"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const articleCollectionName = "article"

type Article struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title    string        `json:"title"`
	UpdateAt time.Time     `json:"updateAt"`
	Author   string        `json:"author"`
	Content  string        `json:"content"`
	Deleted  bool          `json:"deleted"`
}

type Articles []Article

func (a *Article) List(all bool) (error, Articles) {
	db := &DBService
	var query *mgo.Query
	var dbSession *DBSession
	if all {
		query, dbSession = db.Retrieve(articleCollectionName, nil)
	} else {
		query, dbSession = db.Retrieve(articleCollectionName, bson.M{"deleted": false})
	}
	defer dbSession.Close()
	result := make(Articles, 0, 10)
	err := query.Sort("-_id").All(&result)
	return err, result
}

func (a *Article) Save() error {
	db := &DBService
	if a.Id.Valid() {
		res, dbs := db.Update(articleCollectionName, bson.M{"_id": a.Id}, a)
		defer dbs.Close()
		return res
	} else {
		a.Id = bson.NewObjectId()
		res, dbs := db.Create(articleCollectionName, a)
		defer dbs.Close()
		return res
	}
}

func (a *Article) GetOne(articleId string) error {
	db := &DBService
	if !bson.IsObjectIdHex(articleId) {
		return errors.New("paramErr.inValidObjectId/*/" + articleId)
	}
	res, dbs := db.Retrieve(articleCollectionName, bson.M{"_id": bson.ObjectIdHex(articleId)})
	defer dbs.Close()
	return res.One(a)
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
		return errors.New("paramErr.inValidObjectId/*/" + articleId)
	}
	res, dbs := db.Delete(articleCollectionName, bson.M{"_id": bson.ObjectIdHex(articleId)})
	defer dbs.Close()
	return res
}
