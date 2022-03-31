package models

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Authentication map[string]interface{}

type Session struct {
	ID       int       `json:"id" gorm:"column:id"`
	Token    string    `json:"token" gorm:"column:token"`
	UserID   int       `json:"user_id" gorm:"column:user_id"`
	IPAddr   string    `json:"ipaddr" gorm:"column:ipaddr"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
	LoginAt  time.Time `json:"login_at" gorm:"column:login_at"`
}

type ManagerUserDetail struct {
	GroupEnabled bool     `json:"group_enabled" gorm:"column:group_enabled"`
	GroupName    string   `json:"group_name" gorm:"column:group_name"`
	Session      *Session `json:"session" gorm:"-"`
	ManagerUser
}

type SessionParams struct {
	Expire int `json:"expire" gorm:"column:expire"`
}

type TableSessionParams struct {
	Key    string `json:"key" gorm:"column:key"`
	Value  int    `json:"value" gorm:"column:value"`
	Commit string `json:"commit" gorm:"column:commit"`
}

type SessionParamsCache struct {
	lock   sync.Mutex
	Params map[string]int
}

func NewSessionParamsCache() *SessionParamsCache {
	newSPC := SessionParamsCache{
		lock:   sync.Mutex{},
		Params: map[string]int{},
	}
	return &newSPC
}

func (s *SessionParamsCache) Update() {
	s.lock.Lock()
	defer s.lock.Unlock()

	n, bf := getSessionParams()
	if bf != Success {
		return
	}
	s.Params = n
}

func (s *SessionParamsCache) Get(key string) int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Params[key]
}

var SPCache *SessionParamsCache

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
	ID           int       `json:"id" gorm:"column:id"`
	Token        string    `json:"token" gorm:"column:token"`
	IPAddr       string    `json:"ipaddr" gorm:"column:ipaddr"`
	UserID       int       `json:"user_id" gorm:"column:user_id"`
	UpdateAt     time.Time `json:"update_at" gorm:"column:update_at"`
	LoginAt      time.Time `json:"login_at" gorm:"column:login_at"`
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
	if time.Since(usi.UpdateAt).Minutes() > float64(SPCache.Get("session_expire")) { // Session，一小时过期
		config.Log.Errorf("session used: %f, expire: %d", time.Since(usi.UpdateAt).Minutes(), SPCache.Get("session_expire"))
		return ErrLoginExpire
	}
	SyncSessionDate(usi)
	c.Keys = Authentication{
		"userInfo": usi,
	}
	return Success
}

func SyncSessionDate(user *UserSessionInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("session").Where("id=?", user.ID).Update("update_at", time.Now())
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

type ManagerSession struct {
	ID        int       `json:"id" gorm:"column:id"`
	Token     string    `json:"token" gorm:"column:token"`
	IPAddr    string    `json:"ipaddr" gorm:"column:ipaddr"`
	UpdateAt  time.Time `json:"update_at" gorm:"column:update_at"`
	UserID    int       `json:"user_id" gorm:"column:user_id"`
	GroupID   int       `json:"group_id" gorm:"column:group_id"`
	Username  string    `json:"username" gorm:"column:username"`
	Groupname string    `json:"group_name" gorm:"column:group_name"`
}

func GetSession(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	ss := []*ManagerSession{}
	tokens := []*Session{}
	tx := db.Table("session").Find(&tokens)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	count := int64(len(tokens))
	if count == 0 {
		return CalSplitPage(sp, 0, ss), Success
	}
	users := []*ManagerUser{}
	tx = db.Table("manager_user").Find(&users)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	groups := []*ManagerGroup{}
	tx = db.Table("manager_group").Find(&groups)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	usersMap := map[int]*ManagerUser{}
	for _, u := range users {
		usersMap[u.ID] = u
	}
	groupsMap := map[int]*ManagerGroup{}
	for _, g := range groups {
		groupsMap[g.ID] = g
	}
	for _, t := range tokens {
		ms := ManagerSession{
			ID:       t.ID,
			Token:    t.Token,
			IPAddr:   t.IPAddr,
			UpdateAt: t.UpdateAt,
			UserID:   t.UserID,
		}
		if u, ok := usersMap[t.UserID]; ok {
			ms.Username = u.UserName
			if g, ok := groupsMap[u.GroupID]; ok {
				ms.GroupID = g.ID
				ms.Groupname = g.Name
			}
		}
		if sp.Search == "" {
			ss = append(ss, &ms)
			continue
		}
		if strings.Contains(ms.Groupname, sp.Search) ||
			strings.Contains(ms.Username, sp.Search) ||
			strings.Contains(ms.Token, sp.Search) {
			ss = append(ss, &ms)
		}
	}
	return CalSplitPage(sp, count, ss), Success
}

func DelSession(id *OnlyID) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("session").Where("id=?", id.ID).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}

func UpdateSessionParams(sp *SessionParams) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("number_options").Where("`key`='session_expire'").Update(`value`, sp.Expire)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	SPCache.Update()
	return Success
}

func GetSessionParams() ([]*TableSessionParams, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	tsps := []*TableSessionParams{}
	tx := db.Table("number_options").Find(&tsps)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrUpdateData
	}
	return tsps, Success
}

func getSessionParams() (map[string]int, *BriefMessage) {
	data, bf := GetSessionParams()
	if bf != Success {
		return nil, bf
	}
	params := map[string]int{}
	for _, item := range data {
		params[item.Key] = item.Value
	}
	return params, Success
}
