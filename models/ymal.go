package models

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"pro_cfg_manager/config"
	"pro_cfg_manager/utils"
	"sync"

	"gopkg.in/yaml.v3"
)

type MonitorFile struct {
	FileName string
	Data     *MonitroRules
}

type MonitroRules struct {
	Groups []*GroupMember `json:"groups"`
}

type AlertType struct {
	Alert       string            `json:"alert"`
	Expr        string            `json:"expr"`
	For         string            `json:"for"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

type GroupMember struct {
	Name  string       `json:"name"`
	Rules []*AlertType `json:"rules"`
}

type PrometheusRule struct {
	lock sync.Mutex
	data *[]*MonitorFile
}

func NewPR() *PrometheusRule {
	return &PrometheusRule{
		lock: sync.Mutex{},
	}
}

var PR *PrometheusRule

func (pr *PrometheusRule) GetDate() *BriefMessage {
	rulesGroups, bf := GetRulesGroups()
	if bf != Success {
		return bf
	}
	mfs := []*MonitorFile{}
	for _, rg := range rulesGroups {
		mf := &MonitorFile{FileName: rg.Name}
		subGroup, bf := GetSubGroup(rg.ID)
		if bf != Success {
			return bf
		}
		mf.Data = &MonitroRules{
			Groups: []*GroupMember{},
		}
		for _, sg := range subGroup {
			mg := &GroupMember{
				Name:  sg.Name,
				Rules: []*AlertType{},
			}
			mf.Data.Groups = append(mf.Data.Groups, mg)
			mrs, bf := GetMonitorRules(sg.ID, true)
			if bf != Success {
				return bf
			}
			for _, m := range mrs {
				at := &AlertType{
					Alert:       m.Alert,
					Expr:        m.Expr,
					For:         m.For,
					Labels:      map[string]string{},
					Annotations: map[string]string{},
				}
				labels, bf := SearchLabelsByMonitorID(m.ID)
				if bf != Success {
					return bf
				}
				for _, l := range labels {
					at.Labels[l.Key] = utils.ClearCrLR(l.Value)
				}
				annotations, bf := SearchAnnotationsByMonitorID(m.ID)
				if bf != Success {
					return bf
				}
				for _, a := range annotations {
					at.Annotations[a.Key] = utils.ClearCrLR(a.Value)
				}
				mg.Rules = append(mg.Rules, at)
			}
		}
		mfs = append(mfs, mf)
	}
	pr.data = &mfs
	return Success
}

func (pr *PrometheusRule) EmptyData() *BriefMessage {
	if err := filepath.Walk(config.Cfg.PrometheusCfg.RuleConf,
		func(path string, fi os.FileInfo, err error) error {
			if err != nil {
				config.Log.Error(err)
				return nil
			}
			if path == config.Cfg.PrometheusCfg.RuleConf {
				return nil
			}
			fd, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
			if err != nil {
				config.Log.Error(err)
				return err
			}
			if _, err := fd.Seek(4, io.SeekStart); err != nil {
				config.Log.Error(err)
				return err
			}
			err = fd.Truncate(0)
			if err != nil {
				config.Log.Error(err)
				return err
			}
			// return os.RemoveAll(path)
			return nil
		}); err != nil {
		config.Log.Error(err)
	}
	return Success
}

func (pr *PrometheusRule) Write() *BriefMessage {
	if pr.data == nil {
		return ErrDataIsNil
	}
	// 清理文件
	if err := filepath.Walk(config.Cfg.PrometheusCfg.RuleConf,
		func(path string, fi os.FileInfo, err error) error {
			if err != nil {
				config.Log.Error(err)
				return nil
			}
			if path == config.Cfg.PrometheusCfg.RuleConf {
				return nil
			}
			return os.RemoveAll(path)
		}); err != nil {
		config.Log.Error(err)
	}
	for _, mf := range *pr.data {
		yamlData, err := yaml.Marshal(mf.Data)
		if err != nil {
			config.Log.Error(err)
			return ErrDataParse
		}
		absPath := filepath.Join(config.Cfg.PrometheusCfg.RuleConf,
			fmt.Sprintf("%s%s", mf.FileName, MYYaml))
		if err := utils.WIoutilByte(absPath, yamlData); err != nil {
			config.Log.Error(err)
			return ErrDataWriteDisk
		}
	}
	return Success
}

func RulePublish() *BriefMessage {
	if bf := PR.GetDate(); bf != Success {
		return bf
	}
	return PR.Write()
}

func EmptyRulePublish() *BriefMessage {
	return PR.EmptyData()
}
