package memory

import (
	"container/list"
	"github.com/CDog34/GBY/server/libs/session"
	"github.com/pkg/errors"
	"sync"
	"time"
)

type Provider struct {
	lock     sync.Mutex
	sessions map[string]*list.Element
	list     *list.List
}

func (p *Provider) SessionInit(sid string) (session.Session, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newSession := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := p.list.PushBack(newSession)
	p.sessions[sid] = element
	return newSession, nil
}

func (p *Provider) SessionRead(sid string) (session.Session, error) {
	if element, ok := p.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		return nil, errors.New("NoSuchSession")
	}
}
func (p *Provider) SessionDestroy(sid string) error {
	if element, ok := p.sessions[sid]; ok {
		delete(p.sessions, sid)
		p.list.Remove(element)
		return nil
	}
	return nil
}

func (p *Provider) SessionGC(maxLifeTime int64) {
	p.lock.Lock()
	defer p.lock.Unlock()
	for {
		element := p.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxLifeTime) < time.Now().Unix() {
			p.list.Remove(element)
			delete(p.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (p *Provider) SessionUpdate(sid string) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if element, ok := p.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		p.list.MoveToFront(element)
		return nil
	}
	return nil
}

var memoryProvider = &Provider{list: list.New()}
