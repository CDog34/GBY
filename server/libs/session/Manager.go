package session

import (
	"sync"
	"fmt"
	"io"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"
	"time"
)

type SessionManager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
}

func NewSessionManager(providerName, cookieName string, maxLifeTime int64) (*SessionManager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("Unknown Session Provide %s", providerName)
	}
	return &SessionManager{provider:provider, cookieName:cookieName, maxLifeTime:maxLifeTime}, nil
}

func (m *SessionManager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (m *SessionManager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	m.lock.Lock()
	defer m.lock.Unlock()
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		sid := m.sessionId()
		session, _ = m.provider.SessionInit(sid)
		cookie := http.Cookie{Name:m.cookieName, Value:url.QueryEscape(sid), Path:"/", MaxAge:int(m.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = m.provider.SessionRead(sid)
	}
	return
}

func (m *SessionManager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		m.lock.Lock()
		defer m.lock.Unlock()
		m.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := &http.Cookie{Name:m.cookieName, Path:"/", Expires:expiration, MaxAge:-1}
		http.SetCookie(w, cookie)
	}

	return
}

func (m *SessionManager) GC() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.provider.SessionGC(m.maxLifeTime)
	time.AfterFunc(time.Duration(m.maxLifeTime), func() {
		m.GC()
	})
}