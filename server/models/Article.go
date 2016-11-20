package Model

import (
	"time"
	. "github.com/CDog34/GBY/server/services"
	"gopkg.in/mgo.v2/bson"
)

const articleCollectionName = "article"

type Article struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title    string `json:"title"`
	UpdateAt time.Time `json:"updateAt"`
	Author   string `json:"author"`
	Content  string `json:"content"`
}

type Articles []Article

func (a *Article) List() Articles {
	db := &DBService
	query := db.Retrieve(articleCollectionName, nil)
	//defer db.Close()
	result := make(Articles, 0, 10)
	query.All(&result)
	return result
}

func (a *Article) Save() error {
	db := &DBService
	a.Id = bson.NewObjectId()
	return db.Create(articleCollectionName, a)
}

func (a *Article) GetOne(articleId string) error {
	db := &DBService
	return db.Retrieve(articleCollectionName, bson.M{"_id":bson.ObjectIdHex(articleId)}).One(a)
}
