package models

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Authentication map[string]interface{}

type Session struct {
	ID       int       `json:"id" gorm:"column:id"`
	Token    string    `json:"token" gorm:"column:token"`
	UserID   int       `json:"user_id" gorm:"column:user_id"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
}

type ManagerUserDetail struct {
	GroupEnabled bool     `json:"group_enabled" gorm:"column:group_enabled"`
	GroupName    string   `json:"group_name" gorm:"column:group_name"`
	Session      *Session `json:"session" gorm:"-"`
	ManagerUser
}

type SessionCache struct {
	lock           sync.Mutex
	sessionsDetail map[string]*ManagerUserDetail
}

var SSObj *SessionCache

func NewSessionCache() *SessionCache {
	sc := SessionCache{
		lock:           sync.Mutex{},
		sessionsDetail: map[string]*ManagerUserDetail{},
	}
	sc.init()
	return &sc
}

func (s *SessionCache) init() {
	s.lock.Lock()
	defer s.lock.Unlock()

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
		ssMap[s.UserID] = s
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
	// d, ok := s.sessionsDetail[token]
	// if !ok {
	// 	return
	// }
	// go DeleteSession(d.ID)
	// delete(s.sessionsDetail, token)
}

func (s *SessionCache) Flush() {
	s.init()
}

type UserSessionInfo struct {
	Token        string    `json:"token" gorm:"column:token"`
	UserID       int       `json:"user_id" gorm:"column:user_id"`
	UpdateAt     time.Time `json:"update_at" gorm:"column:update_at"`
	Username     string    `json:"username" gorm:"column:username"`
	NiceName     string    `json:"nice_name" gorm:"column:nice_name"`
	GroupID      int       `json:"group_id" gorm:"column:group_id"`
	Phone        string    `json:"phone" gorm:"column:phone"`
	UserEnabled  bool      `json:"user_enabled" gorm:"column:user_enabled"`
	GroupEnabled bool      `json:"group_enabled" gorm:"column:group_enabled"`
}

func GetUserInformation(token string) (*UserSessionInfo, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	userSession := UserSessionInfo{}
	tx := db.Table("session").
		Select("session.*, manager_user.username, manager_user.group_id, manager_user.phone, manager_user.enabled AS user_enabled, manager_group.name, manager_group.enabled AS group_enabled").
		Joins("LEFT JOIN manager_user ON session.user_id=manager_user.id").
		Joins("LEFT JOIN manager_group ON manager_group.id=manager_user.group_id").
		Where("session.token=?", token).Find(&userSession)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	if tx.RowsAffected == 0 {
		return nil, ErrTokenIsNull
	}
	return &userSession, Success
}

func CheckUserSession(c *gin.Context, token string) *BriefMessage {
	usi, bf := GetUserInformation(token)
	if bf != Success {
		return ErrTokenNoFound
	}
	c.Keys = Authentication{
		"userInfo": usi,
	}
	return Success
}
