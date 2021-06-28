package models

import (
	"fmt"
	"path/filepath"
	"pro_cfg_manager/config"
	"pro_cfg_manager/utils"
	"sync"

	"gopkg.in/yaml.v2"
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

var PR = NewPR()

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
					at.Labels[l.Key] = l.Value
				}
				annotations, bf := SearchAnnotationsByMonitorID(m.ID)
				if bf != Success {
					return bf
				}
				for _, a := range annotations {
					at.Annotations[a.Key] = a.Value
				}
				mg.Rules = append(mg.Rules, at)
			}
		}
		mfs = append(mfs, mf)
	}
	pr.data = &mfs
	return Success
}

func (pr *PrometheusRule) Write() *BriefMessage {
	if pr.data == nil {
		return ErrDataIsNil
	}
	for _, mf := range *pr.data {
		yamlData, err := yaml.Marshal(mf.Data)
		if err != nil {
			config.Log.Error(err)
			return ErrDataParse
		}
		absPath := filepath.Join(config.Cfg.PrometheusCfg.RuleConf, fmt.Sprintf("%s%s", mf.FileName, MYYaml))
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
