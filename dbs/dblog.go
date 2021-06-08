package dbs

import (
	"context"
	"fmt"
	"pro_cfg_manager/config"
	"strings"
	"time"

	"gorm.io/gorm/logger"
)

type dbLog struct{}

func newLogrusForDB() logger.Interface {
	newLog := logger.New(
		config.Log,
		// log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // 彩色打印
		},
	)
	// newLog.LogMode(logger.Info)
	return newLog
}

func newDBLog() logger.Interface {
	var li logger.Interface
	li = new(dbLog)
	return li
}

func (d *dbLog) LogMode(logger.LogLevel) logger.Interface {
	// config.Log.Printf("db logger level: %s", l)
	return d
}

func (d *dbLog) Info(ctx context.Context, s string, is ...interface{}) {
	t := fmt.Sprintf(s, is...)
	config.Log.Info(strings.TrimSpace(t))
}

func (d *dbLog) Warn(ctx context.Context, s string, is ...interface{}) {
	t := fmt.Sprintf(s, is...)
	config.Log.Warnf(strings.TrimSpace(t))
}

func (d *dbLog) Error(ctx context.Context, s string, is ...interface{}) {
	t := fmt.Sprintf(s, is...)
	config.Log.Error(strings.TrimSpace(t))
}

func (d *dbLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	r, i := fc()
	s := fmt.Sprintf("begin: %s, %s, %d, %s", begin, r, i, err)
	config.Log.Trace(s)
}
