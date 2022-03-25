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
	RootDir string `json:"dir" yaml:"dir"`
	// TmplFile    string `json:"tmpl" yaml:"tmpl"`
	Api         string `json:"api" yaml:"api"`
	OpenAddress string `json:"open_address" yaml:"open_address"`
	Username    string `json:"username" yaml:"username"`
	Password    string `json:"password" yaml:"password"`
	// TmplContext string `json:"tmpl_context" yaml:"-"`
	MainConf string `json:"main_conf" yaml:"-"`
	RuleConf string `json:"rule_conf" yaml:"-"`
	Conf     string `json:"conf" yaml:"-"`
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
	Position      string        `json:"position" yaml:"position"`
	RuntimeParam  runtimeParam  `json:"-" yaml:"-"`
}

// Cfg 全局的Config配置，解析dns.yaml的结果
var Cfg *Config

// Version 软件版本
var Version = "1.0.3 bate"

// SubDir for config
var SubDir = "conf.d"
var RuleDir = "rules"

//
var PrometheusConfigName = "prometheus.yml"

//
var configFileName = "config.yml"
