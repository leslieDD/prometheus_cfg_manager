package dbs

import (
	"pro_cfg_manager/config"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBObj DBObj
var DBObj = NewInitDB()

// NewInitDB NewInitDB
func NewInitDB() *InitDB {
	idb := &InitDB{}
	idb.lock = sync.Mutex{}
	return idb
}

// InitDB 初始化数据库的连接
type InitDB struct {
	// DBConn 连接实例
	dbConn *gorm.DB
	lock   sync.Mutex
}

// Init Init
func (i *InitDB) Init() {
	i.initMysql()
}

func (i *InitDB) initMysql() (done bool) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	db, err := gorm.Open(mysql.Open(config.Cfg.Mysql.URL), &gorm.Config{
		// Logger: newLogrusForDB().LogMode(logger.Error),
		Logger: newLogrusForDB(),
	})
	if err != nil {
		config.Log.Error(err)
		return
	}
	// db.Logger.LogMode(logger.Info)
	i.dbConn = db
	return true
}

func (i *InitDB) initMongo() {}

func (i *InitDB) initSqlite() {}

func (i *InitDB) initClickHouse() {}

func (i *InitDB) initTiDB() {}

// GetGoRM 得到数据库连接实例
func (i *InitDB) GetGoRM() *gorm.DB {
	if i.dbConn == nil {
		if !i.initMysql() {
			return nil
		}
	}
	// i.dbConn.Session()
	return i.dbConn
}
