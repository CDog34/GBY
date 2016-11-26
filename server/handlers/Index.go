package handlers

import (
	"github.com/CDog34/GBY/server/services"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) (error, interface{}) {
	session, err := services.SessionMgr.SessionStart(w, r, false)
	if err != nil {
		return errors.New("auth.needSession"), nil
	}
	count := 0
	if num := session.Get("count"); num != nil {
		count = num.(int)
	}
	session.Set("count", count+1)
	return nil, map[string]interface{}{
		"system": "ok",
		"time":   time.Now(),
		"count":  session.Get("count"),
		"sid":    session.SessionID(),
	}
}
