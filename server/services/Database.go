package services

import (
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"github.com/CDog34/GBY/server/configs"
)

type Database struct {
	url     string
	dbName  string
	session *mgo.Session
}
type DBSession struct {
	session *mgo.Session
}

func (dbs *DBSession) Close() {
	dbs.session.Close()
}

var DBService = Database{
	url:    config.Get("dbUrl").(string),
	dbName: config.Get("dbName").(string),
}

func hasPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f()
	return
}
func (db *Database) getSession() *DBSession {
	if db.session != nil {
		isBroken := hasPanic(func() {
			err := db.session.Ping()
			if err != nil {
				panic("SessionError")
			}
		})
		if !isBroken {
			return &DBSession{db.session.Copy()}
		}
	}
	dbSession, err := mgo.Dial(db.url)
	if err != nil {
		panic(err)
	}
	db.session = dbSession
	return &DBSession{db.session.Copy()}
}

func (db *Database) selectCollection(collection string) (*mgo.Collection, *DBSession) {
	sessionWrapper := db.getSession()
	dbSession := sessionWrapper.session
	dbInstance := dbSession.DB(db.dbName)
	dbCollection := dbInstance.C(collection)
	return dbCollection, sessionWrapper
}

func (db *Database) Close() {
	db.session.Close()
}

func (db *Database) Create(collection string, data interface{}) (error, *DBSession) {
	dbCollection, dbSession := db.selectCollection(collection)
	err := dbCollection.Insert(data)
	return err, dbSession
}

func (db *Database) Retrieve(collection string, query map[string]interface{}) (*mgo.Query, *DBSession) {
	dbCollection, dbSession := db.selectCollection(collection)
	result := dbCollection.Find(query)
	return result, dbSession
}

func (db *Database) Update(collection string, query map[string]interface{}, data interface{}) (error, *DBSession) {
	dbCollection, dbSession := db.selectCollection(collection)
	result := dbCollection.Update(query, data)
	return result, dbSession
}

func (db *Database) Delete(collection string, query map[string]interface{}) (error, *DBSession) {
	dbCollection, dbSession := db.selectCollection(collection)
	result := dbCollection.Remove(query)
	return result, dbSession
}
