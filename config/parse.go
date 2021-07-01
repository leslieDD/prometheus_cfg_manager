package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"pro_cfg_manager/utils"

	"gopkg.in/yaml.v2"
)

func init() {
	var err error
	fmt.Printf("version: %s\n", Version)
	rootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	if !utils.FileExists(configFileName) {
		configFileName = filepath.Join(rootDir, configFileName)
	}
	Cfg, err = ParseConfig(configFileName, rootDir)
	if err != nil {
		log.Fatal(err)
	}
}

// ParseConfig ParseConfig
func ParseConfig(cfgPath string, rootDir string) (*Config, error) {
	fd, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	config.RuntimeParam.RootDir = rootDir
	err = yaml.NewDecoder(fd).Decode(config)
	if err != nil {
		return nil, err
	}
	if err := config.verification(); err != nil {
		return nil, err
	}
	return config, nil
}

func (cf *Config) verification() error {
	// todo: 做字段的检测
	if cf.Logger.Path == "" {
		cf.Logger.Path = "logs"
	}
	cf.Logger.Path = filepath.Join(cf.RuntimeParam.RootDir, cf.Logger.Path)
	_, err := os.Stat(cf.Logger.Path)
	if err != nil {
		if mkErr := os.Mkdir(cf.Logger.Path, 0644); mkErr != nil {
			return mkErr
		}
	}
	if exist, _ := utils.PathExists(cf.PrometheusCfg.RootDir); !exist {
		if mkErr := os.Mkdir(cf.PrometheusCfg.RootDir, 0644); mkErr != nil {
			return mkErr
		}
	}
	cf.PrometheusCfg.Conf = filepath.Join(cf.PrometheusCfg.RootDir, SubDir)
	if exist, _ := utils.PathExists(cf.PrometheusCfg.Conf); !exist {
		if mkErr := os.Mkdir(cf.PrometheusCfg.Conf, 0644); mkErr != nil {
			return mkErr
		}
	}
	cf.PrometheusCfg.RuleConf = filepath.Join(cf.PrometheusCfg.RootDir, RuleDir)
	if exist, _ := utils.PathExists(cf.PrometheusCfg.RuleConf); !exist {
		if mkErr := os.Mkdir(cf.PrometheusCfg.RuleConf, 0644); mkErr != nil {
			return mkErr
		}
	}
	cf.PrometheusCfg.MainConf = filepath.Join(cf.PrometheusCfg.RootDir, PrometheusConfigName)
	// tmplTxt, err := utils.RIoutil(cf.PrometheusCfg.TmplFile)
	// if err != nil {
	// 	cf.PrometheusCfg.TmplFile = filepath.Join(cf.RuntimeParam.RootDir, cf.PrometheusCfg.TmplFile)
	// 	tmplTxt, err = utils.RIoutil(cf.PrometheusCfg.TmplFile)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// cf.PrometheusCfg.TmplContext = string(tmplTxt)
	// cf.PrometheusCfg.TmplFile = filepath.Join(cf.RuntimeParam.RootDir, cf.PrometheusCfg.TmplFile)
	return nil
}
