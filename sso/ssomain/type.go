package ssomain

import "sync"

type serverinfo struct {
	token    []string
	userinfo string
	appid    []string
}
type serverLogined struct {
	lock   sync.RWMutex
	server map[string]*serverinfo
}

func (s *serverLogined) AddToken(sessionid string, token string) {
	if s.server == nil {
		s.server = make(map[string]*serverinfo, 10)
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.server[sessionid] == nil {
		s.server[sessionid] = &serverinfo{}
	}
	s.server[sessionid].token = append(s.server[sessionid].token, token)
}
func (s *serverLogined) AddAppid(sessionid string, appid string) {
	if s.server == nil {
		s.server = make(map[string]*serverinfo, 10)
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.server[sessionid] == nil {
		s.server[sessionid] = &serverinfo{}
	}
	s.server[sessionid].appid = append(s.server[sessionid].appid, appid)
}
func (s *serverLogined) AddUser(sessionid string, username string) {
	if s.server == nil {
		s.server = make(map[string]*serverinfo, 10)
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.server[sessionid] == nil {
		s.server[sessionid] = &serverinfo{}
	}
	s.server[sessionid].userinfo = username
}
func (s *serverLogined) Get(sessionid string) *serverinfo {
	s.lock.RLock()
	defer s.lock.RUnlock()
	tokenlist, ok := s.server[sessionid]
	if !ok {
		return nil
	}
	return tokenlist
}
func (s *serverLogined) Del(sessionid string) {
	s.lock.Lock()
	delete(s.server, sessionid)
	s.lock.Unlock()

	return
}
