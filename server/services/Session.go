package services

import (
	"github.com/CDog34/GBY/server/libs/session"
	_ "github.com/CDog34/GBY/server/libs/session/providers/memory"
)

var SessionMgr *session.SessionManager

func init() {
	SessionMgr, _ = session.NewSessionManager("memory", "moeMisaka", 3600)
	go SessionMgr.GC()
}
