package alertmgr

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"strings"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type ChartCron struct {
	lock      sync.Mutex
	RulesInfo []*CronRule
	doAction  chan struct{}

	cObj    *cron.Cron
	entryID []cron.EntryID
}

var ChartCronObj *ChartCron

func NewChartCron() *ChartCron {
	return &ChartCron{
		lock:     sync.Mutex{},
		doAction: make(chan struct{}, 1),
		entryID:  []cron.EntryID{},
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
		// 只取前6个字段
		ruleFiles := strings.Fields(rule.ExecCycle)
		var execRule string
		// config.Log.Print("start rule: ", rule.ExecCycle, rule.Name)
		if len(ruleFiles) >= 6 {
			execRule = strings.Join(ruleFiles[0:5], " ")
		} else {
			execRule = rule.ExecCycle
		}
		c.DoDoDo(execRule, rule)
	}
	c.cObj.Start()

	c.lock.Unlock()
	config.Log.Warnf("start cron done")
}

// 单独写，是因为上面写在一个函数里面的时候，会出现参数函数里面，rule只会用最后一个
func (c *ChartCron) DoDoDo(execRule string, rule *CronRule) {
	/*
		# 文件格式說明
		# ┌──分鐘（0 - 59）
		# │  ┌──小時（0 - 23）
		# │  │  ┌──日（1 - 31）
		# │  │  │  ┌─月（1 - 12）
		# │  │  │  │  ┌─星期（0 - 6，表示从周日到周六）
		# │  │  │  │  │
		# *  *  *  *  * 被執行的命令
	*/
	eid, err := c.cObj.AddFunc(execRule, func() {
		c.Work(rule)
	})
	if err != nil {
		config.Log.Error(err)
		return
	}
	c.entryID = append(c.entryID, eid)
}

func (c *ChartCron) Work(rule *CronRule) {
	defer func() {
		if r := recover(); r != nil {
			config.Log.Error("Recovered in Work", r)
		}
	}()
	ic := ImageContent{}
	image, err := ChartLine(rule)
	if err != nil {
		return
	}
	ic.Img = image
	sendEmail(&ic)
}
