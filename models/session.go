package models

import (
	"sync"
	"time"
)

type Authentication map[string]interface{}

type Session struct {
	ID       int       `json:"id" gorm:"id"`
	Token    string    `json:"token" gorm:"token"`
	UserID   int       `json:"user_id" gorm:"user_id"`
	UpdateAt time.Time `json:"update_at" gorm:"update_at"`
}

type ManagerUserDetail struct {
	GroupEnabled bool     `json:"group_enabled" gorm:"group_enabled"`
	GroupName    string   `json:"group_name" gorm:"group_name"`
	Session      *Session `json:"session" gorm:"-"`
	ManagerUser
}

type SessionCache struct {
	lock           sync.Mutex
	sessionsDetail map[string]*ManagerUserDetail
}

var SSObj = NewSessionCache()

func NewSessionCache() *SessionCache {
	sc := SessionCache{
		lock:           sync.Mutex{},
		sessionsDetail: map[string]*ManagerUserDetail{},
	}
	sc.init()
	return &sc
}

func (s *SessionCache) init() {
	us, bf := LoadUserEnabled()
	if bf != Success {
		return
	}
	ids := []int{}
	for _, u := range us {
		if !u.Enabled || !u.GroupEnabled {
			continue
		}
		ids = append(ids, u.ID)
	}
	ss, bf := LoadSession(ids)
	if bf != Success {
		return
	}
	ssMap := map[int]*Session{}
	for _, s := range ss {
		ssMap[s.ID] = s
	}
	for _, u := range us {
		ss, ok := ssMap[u.ID]
		if !ok {
			continue
		}
		if time.Now().Sub(ss.UpdateAt).Hours() > 12 {
			continue
		}
		u.Session = ss
		s.sessionsDetail[u.Session.Token] = u
	}
}

func (s *SessionCache) Get(token string) *ManagerUserDetail {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.sessionsDetail[token]
}

func (s *SessionCache) Set(token string, ss *ManagerUserDetail) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.sessionsDetail[token] = ss
}

func (s *SessionCache) Del(token string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	d, ok := s.sessionsDetail[token]
	if !ok {
		return
	}
	go DeleteSession(d.ID)
	delete(s.sessionsDetail, token)
}
