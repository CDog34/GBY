package handlers

import (
	"errors"
	"fmt"
	"github.com/CDog34/GBY/server/configs"
	"github.com/CDog34/GBY/server/models"
	"github.com/CDog34/GBY/server/services"
	"net/http"
	"time"
)

var UserPostRules = services.FieldRules{
	"email": services.FieldRule{
		"required": true,
		"type":     "string",
	},
	"password": services.FieldRule{
		"required": true,
		"type":     "string",
	},
	"passwordConfirm": services.FieldRule{
		"required": false,
		"type":     "string",
	},
}

func AuthLogin(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	sess, serr := services.SessionMgr.SessionStart(w, r, false)
	if serr == nil && sess.Get("user") != nil {
		return nil, nil
	}
	params := services.PostParams{Request: r, Rules: UserPostRules}
	if err := params.Valid(); err != nil {
		return err, nil
	}
	user := Model.User{}
	user.GetOneByEmail(params.GetString("email"))
	if !user.Id.Valid() {
		return fmt.Errorf("auth.authFail/*/%s", params.GetString("email")), nil
	}
	if !user.CheckPassword(params.GetString("password")) {
		return fmt.Errorf("auth.authFail/*/%s", params.GetString("email")+"pwd"), nil
	}
	sess, _ = services.SessionMgr.SessionStart(w, r, true)
	sess.Set("user", user)
	return nil, map[string]interface{}{
		"name":  config.Get("sessionName"),
		"value": sess.SessionID(),
		"user":  user,
	}

}

func AuthNewUser(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	user := Model.User{}
	if _, userList := user.List(); len(userList) > 0 {
		return errors.New("auth.notAllow"), nil
	}

	params := services.PostParams{Request: r, Rules: UserPostRules}
	if err := params.Valid(); err != nil {
		return err, nil
	}
	if params.GetString("password") != params.GetString("passwordConfirm") {
		return errors.New("auth.passwordNotMatch"), nil
	}
	user.GetOneByEmail(params.GetString("email"))
	if user.Id.Valid() {
		return fmt.Errorf("auth.emailAlreadyExist/*/%s", params.GetString("email")), nil
	}
	user.Email = params.GetString("email")
	user.GeneratePassword(params.GetString("password"))
	user.Active = true
	user.Role = 100
	user.UpdateAt = time.Now()
	err := user.Save()
	return err, user
}
