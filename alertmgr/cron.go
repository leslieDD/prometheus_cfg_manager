package alertmgr

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type ChartCron struct {
	lock      sync.Mutex
	RulesInfo []*CronRule
	doAction  chan struct{}

	cObj *cron.Cron
}

var ChartCronObj *ChartCron

func NewChartCron() *ChartCron {
	return &ChartCron{
		lock:     sync.Mutex{},
		doAction: make(chan struct{}, 1),
	}
}

func (c *ChartCron) Run() {
	c.LoadRule()
	go c.DoCron()
	go c.RecycleLoadRule()
}

func (c *ChartCron) LoadRule() {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error("get db session err")
		return
	}
	rules := []*CronRule{}
	tx := db.Raw(`SELECT crontab.*, crontab_api.api, crontab_api.name as api_name FROM crontab 
	LEFT join crontab_api 
	on crontab.api_id=crontab_api.id
	WHERE crontab.enabled=1
	`).Find(&rules)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	config.Log.Warnf("load cron rules: %d", len(rules))
	c.RulesInfo = rules
}

func (c *ChartCron) RecycleLoadRule() {
	for {
		select {
		case <-c.doAction:
			c.LoadRule()
			c.DoCron()
		}
	}
}

func (c *ChartCron) NoticeReload() {
	c.doAction <- struct{}{}
}

func (c *ChartCron) DoCron() {
	if c.cObj != nil {
		c.cObj.Stop()
	}
	nyc, _ := time.LoadLocation("Asia/Shanghai")
	c.cObj = cron.New(cron.WithLocation(nyc))

	c.lock.Lock()

	for _, rule := range c.RulesInfo {
		c.cObj.AddFunc(rule.Name, func() {
			ChartLine(rule)
			// config.Log.Warnf(rule.ExecCycle)
		})
	}
	c.cObj.Start()

	c.lock.Unlock()
}
