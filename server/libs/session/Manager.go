package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type SessionManager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
	useHeader   bool
}

func NewSessionManager(providerName, cookieName string, maxLifeTime int64, useHeader bool) (*SessionManager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("Unknown Session Provide %s", providerName)
	}
	return &SessionManager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime, useHeader: useHeader}, nil
}

func (m *SessionManager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
func (m *SessionManager) getSid(w http.ResponseWriter, r *http.Request) (sid string, err error) {
	var value string
	if m.useHeader {
		value = r.Header.Get(m.cookieName)
	}
	if value == "" {
		var cookie *http.Cookie
		cookie, err = r.Cookie(m.cookieName)
		if err != nil {
			return
		}
		value = cookie.Value
	}
	if value == "" {
		err = errors.New("NoSession")
	} else {
		sid, err = url.QueryUnescape(value)
	}
	return
}

func (m *SessionManager) setSid(sid string, w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: m.cookieName, Value: url.QueryEscape(sid), Path: "/", MaxAge: int(m.maxLifeTime)}
	http.SetCookie(w, &cookie)
	if m.useHeader {
		w.Header().Set(m.cookieName, sid)
	}
}

func (m *SessionManager) SessionStart(w http.ResponseWriter, r *http.Request, createNew bool) (session Session, err error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	sid, sidErr := m.getSid(w, r)
	if sidErr != nil {
		if createNew {
			sid = m.sessionId()
			session, _ = m.provider.SessionInit(sid)
			m.setSid(sid, w, r)
		} else {
			err = errors.New("NoSession")
		}
	} else {
		session, err = m.provider.SessionRead(sid)
		if err != nil && createNew {
			session, _ = m.provider.SessionInit(sid)
			err = nil
		}
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
		cookie := &http.Cookie{Name: m.cookieName, Path: "/", Expires: expiration, MaxAge: -1}
		http.SetCookie(w, cookie)
	}

	return
}

func (m *SessionManager) GC() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.provider.SessionGC(m.maxLifeTime)
	time.AfterFunc(time.Duration(m.maxLifeTime*1000*1000), func() {
		m.GC()
	})
}
