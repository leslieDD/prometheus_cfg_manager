package config

type app struct {
	Name   string `json:"name" yaml:"name"`
	Listen string `json:"listen" yaml:"listen"`
}

type logger struct {
	Path  string `json:"path" yaml:"path"`
	Level string `json:"level" yaml:"level"`
}

type mysql struct {
	URL string `json:"url" yaml:"url"`
}

type prometheusCfg struct {
	Dir string `json:"dir" yaml:"dir"`
}

type runtimeParam struct {
	RootDir string `json:"-" yaml:"-"` // 此软件运行后的工作目录
}

// Config 全局的Config配置，解析dns.yaml的结果
type Config struct {
	App           app           `json:"app" yaml:"app"`
	Logger        logger        `json:"logger" yaml:"logger"`
	Mysql         mysql         `json:"mysql" yaml:"mysql"`
	PrometheusCfg prometheusCfg `json:"prometheus_cfg" yaml:"prometheus_cfg"`
	RuntimeParam  runtimeParam  `json:"-" yaml:"-"`
}

// Cfg 全局的Config配置，解析dns.yaml的结果
var Cfg *Config

// Version 软件版本
var Version = "0.0.5"

//
var configFileName = "pro_cfg_manager.yml"
