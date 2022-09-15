package alertmgr

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/models"
	"strings"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type ChartCron struct {
	lock      sync.Mutex
	RulesInfo map[int]*CronRule
	doAction  chan struct{}

	cObj    *cron.Cron
	entryID map[int]cron.EntryID
}

var ChartCronObj *ChartCron

func NewChartCron() *ChartCron {
	taskRunStatusObj = &taskRunStatus{
		lock:   sync.Mutex{},
		status: map[int]*taskRst{},
	}
	return &ChartCron{
		lock:      sync.Mutex{},
		doAction:  make(chan struct{}, 1),
		entryID:   map[int]cron.EntryID{},
		RulesInfo: map[int]*CronRule{},
	}
}

func (c *ChartCron) Run() {
	c.RulesInfo = c.LoadRule()
	go c.DoCron()
	go c.RecycleLoadRule()
}

func (c *ChartCron) LoadRule() map[int]*CronRule {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error("get db session err")
		return nil
	}
	rules := []*CronRule{}
	tx := db.Raw(`SELECT crontab.*, crontab_api.api, crontab_api.name as api_name FROM crontab 
	LEFT join crontab_api 
	on crontab.api_id=crontab_api.id
	WHERE crontab.enabled=1
	`).Find(&rules)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil
	}
	tasks := map[int]*CronRule{}
	for _, rule := range rules {
		tasks[rule.ID] = rule
	}
	config.Log.Warnf("load cron rules: %d", len(tasks))
	return tasks
}

func (c *ChartCron) LoadRuleManual() {
	rulesInfo := c.LoadRule()
	c.lock.Lock()
	defer c.lock.Unlock()
	// 移除所有已经不存在的任务
	for ruleID, taskID := range c.entryID {
		_, ok := rulesInfo[ruleID]
		if !ok {
			c.cObj.Remove(taskID)
			config.Log.Warnf("remove cron task: %d", ruleID)
			continue
		}
	}
	for ruleID, newRule := range rulesInfo {
		oldRule := c.RulesInfo[ruleID]
		if oldRule != nil {
			if !oldRule.UpdateAt.Equal(newRule.UpdateAt) {
				config.Log.Warnf("stop cron task: %d, will start soon", ruleID)
				c.cObj.Remove(c.entryID[ruleID])
			} else {
				continue
			}
		}
		execRule := c.ConvertSpec(newRule)
		c.DoDoDo(execRule, newRule)
		taskRunStatusObj.Reset(newRule.ID)
	}
	c.RulesInfo = rulesInfo
}

func (c *ChartCron) RecycleLoadRule() {
	for {
		select {
		case <-c.doAction:
			c.LoadRuleManual()
		}
	}
}

func (c *ChartCron) NoticeReload() {
	c.doAction <- struct{}{}
}

func (c *ChartCron) DoCron() {
	nyc, _ := time.LoadLocation("Asia/Shanghai")
	c.cObj = cron.New(cron.WithLocation(nyc))

	c.lock.Lock()
	defer c.lock.Unlock()

	for _, rule := range c.RulesInfo {
		execRule := c.ConvertSpec(rule)
		c.DoDoDo(execRule, rule)
		taskRunStatusObj.Reset(rule.ID)
	}
	c.cObj.Start()
	config.Log.Warnf("start cron done")
}

func (c *ChartCron) ConvertSpec(rule *CronRule) string {
	ruleFiles := strings.Fields(rule.ExecCycle)
	var execRule string
	if len(ruleFiles) >= 5 {
		execRule = strings.Join(ruleFiles[0:5], " ")
	} else {
		execRule = rule.ExecCycle
	}
	return execRule
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
	config.Log.Warnf("start cron task: %d, %d", rule.ID, eid)
	c.entryID[rule.ID] = eid
}

func (c *ChartCron) Work(rule *CronRule) {
	defer func() {
		if r := recover(); r != nil {
			config.Log.Error("Recovered in Work", r)
			taskRunStatusObj.AddFail(rule.ID)
		}
	}()
	ic := ImageContent{
		Title:  rule.Name,
		Width:  1000,
		Height: 600,
	}
	image, err := ChartLine(rule)
	if err != nil {
		taskRunStatusObj.AddFail(rule.ID)
		return
	}
	ic.Image = image
	if sendEmail(&ic) {
		taskRunStatusObj.AddSuccess(rule.ID)
	} else {
		taskRunStatusObj.AddFail(rule.ID)
	}
}

// alertmgr和models冲突，就这么搞了
func MergeRuleStatus(data *models.ResSplitPage) {
	if data == nil {
		return
	}
	rules := data.Data.([]*models.CrontabSplit)
	status := taskRunStatusObj.SyncStatus()
	for _, rule := range rules {
		obj, ok := status[rule.ID]
		if !ok {
			continue
		}
		rule.RunTimes = obj.RunTimes
		rule.SuccessTimes = obj.SuccessTimes
		rule.FailTimes = obj.FailTimes
	}
	data.Data = rules
}
