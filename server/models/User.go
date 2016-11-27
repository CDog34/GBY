package Model

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"errors"
	. "github.com/CDog34/GBY/server/services"
	"gopkg.in/mgo.v2/bson"
	"io"
	"time"
)

const userCollectionName = "user"
const saltLength = 32

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email    string        `json:"email"`
	Password []byte        `json:"-" bson:"password"`
	Active   bool          `json:"active"`
	NickName string        `json:"nickName"`
	Role     int           `json:"role"`
	UpdateAt time.Time     `json:"updateAt"`
}
type Users []User

func (u *User) List() (error, Users) {
	db := &DBService
	query, dbSession := db.Retrieve(userCollectionName, nil)
	defer dbSession.Close()
	result := make(Users, 0, 10)
	err := query.All(&result)
	return err, result
}

func (u *User) Save() error {
	db := &DBService
	if u.Id.Valid() {
		res, dbs := db.Update(userCollectionName, bson.M{"_id": u.Id}, u)
		defer dbs.Close()
		return res
	} else {
		u.Id = bson.NewObjectId()
		res, dbs := db.Create(userCollectionName, u)
		defer dbs.Close()
		return res
	}
}

func (u *User) GetOne(userId string) error {
	db := &DBService
	if !bson.IsObjectIdHex(userId) {
		return errors.New("paramErr.inValidObjectId/*/" + userId)
	}
	res, dbs := db.Retrieve(userCollectionName, bson.M{"_id": bson.ObjectIdHex(userId)})
	defer dbs.Close()
	return res.One(u)
}
func (u *User) GetOneByEmail(email string) error {
	db := &DBService
	res, dbs := db.Retrieve(userCollectionName, bson.M{"email": email})
	defer dbs.Close()
	return res.One(u)
}
func (u *User) HardDelete(userId string) error {
	db := &DBService
	if !bson.IsObjectIdHex(userId) {
		return errors.New("paramErr.inValidObjectId/*/" + userId)
	}
	res, dbs := db.Delete(userCollectionName, bson.M{"_id": bson.ObjectIdHex(userId)})
	defer dbs.Close()
	return res
}

func (u *User) generateSalt() []byte {
	b := make([]byte, saltLength)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return []byte("")
	}
	return b
}

func (u *User) GeneratePassword(password string) {
	passwordByte := []byte(password)
	salt := u.generateSalt()
	material := bytes.Join([][]byte{salt, passwordByte}, nil)
	secret := sha1.Sum(material)
	u.Password = append(salt, secret[:]...)
}

func (u *User) CheckPassword(inputPassword string) bool {
	salt := u.Password[:saltLength]
	inputPasswordByte := []byte(inputPassword)
	material := bytes.Join([][]byte{salt, inputPasswordByte}, nil)
	inputSecret := sha1.Sum(material)
	return bytes.Equal(u.Password[saltLength:], inputSecret[:])
}
