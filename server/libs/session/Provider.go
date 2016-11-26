package session

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(MaxLifeTime int64)
}

var providers = make(map[string]Provider)

func Register(name string, driver Provider) {
	if driver == nil {
		panic("Register provider is nil");
	}
	if _, dup := providers[name]; dup {
		panic("Rgister Called Twice for " + name)
	}
	providers[name] = driver
}