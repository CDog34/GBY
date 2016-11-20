package services

import (
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

type Database struct {
	url     string
	dbName  string
	session *mgo.Session
}

var DBService = Database{
	url:"localhost",
	dbName:"GBY_DEV",
};

func hasPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f()
	return
}
func (db *Database) getSession() *mgo.Session {
	if db.session != nil {
		isBroken := hasPanic(func() {
			err := db.session.Ping()
			if err != nil {
				panic("SessionError")
			}
		})
		if !isBroken {
			return db.session
		}
	}
	session, err := mgo.Dial(db.url)
	if err != nil {
		panic(err)
	}
	db.session = session
	return session
}

func (db *Database) selectCollection(collection string) *mgo.Collection {
	dbSession := db.getSession()
	dbInstance := dbSession.DB(db.dbName)
	dbCollection := dbInstance.C(collection)
	return dbCollection
}

func (db *Database) Close() {
	db.session.Close()
}

func (db *Database) Create(collection string, data interface{}) error {
	dbCollection := db.selectCollection(collection)
	err := dbCollection.Insert(data)
	return err
}

func (db *Database) Retrieve(collection string, query map[string]interface{}) *mgo.Query {
	dbCollection := db.selectCollection(collection)
	result := dbCollection.Find(query)
	return result
}

func (db *Database) Update(collection string, query map[string]interface{}, data interface{}) error {
	dbCollection := db.selectCollection(collection)
	result := dbCollection.Update(query, data)
	return result
}