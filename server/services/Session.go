package services

import (
	"github.com/CDog34/GBY/server/configs"
	"github.com/CDog34/GBY/server/libs/session"
	_ "github.com/CDog34/GBY/server/libs/session/providers/memory"
)

var SessionMgr *session.SessionManager

func init() {
	SessionMgr, _ = session.NewSessionManager(config.Get("sessionProvider").(string), config.Get("sessionName").(string), int64(config.Get("sessionLifeTime").(int)), true)
	go SessionMgr.GC()
}
