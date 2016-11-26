package memory

import (
	"time"
	"container/list"
	"github.com/CDog34/GBY/server/libs/session"
)

type SessionStore struct {
	sid          string
	timeAccessed time.Time
	value        map[interface{}]interface{}
}

func init() {
	memoryProvider.sessions = make(map[string]*list.Element, 0)
	session.Register("memory", memoryProvider)
}

func (st *SessionStore)Set(key, value interface{}) error {
	st.value[key] = value
	memoryProvider.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore)Get(key interface{}) interface{} {
	memoryProvider.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
}

func (st *SessionStore)Delete(key interface{}) error {
	delete(st.value, key)
	memoryProvider.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) SessionID() string {
	return st.sid
}


