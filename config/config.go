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

type gps struct {
	Geo       string `json:"geo" yaml:"geo"`
	IP2region string `json:"ip2region" yaml:"ip2region"`
}

type mail struct {
	Sender   string `json:"sender" yaml:"sender"`
	Password string `json:"password" yaml:"password"`
	SMTP     string `json:"smtp" yaml:"smtp"`
	Port     int    `json:"port" yaml:"port"`
}

type prometheusCfg struct {
	RootDir string `json:"dir" yaml:"dir"`
	// TmplFile    string `json:"tmpl" yaml:"tmpl"`
	Api         string `json:"api" yaml:"api"`
	OpenAddress string `json:"open_address" yaml:"open_address"`
	Username    string `json:"username" yaml:"username"`
	Password    string `json:"password" yaml:"password"`
	// TmplContext string `json:"tmpl_context" yaml:"-"`
	MainConf     string `json:"main_conf" yaml:"-"`
	RuleConf     string `json:"rule_conf" yaml:"-"`
	Conf         string `json:"conf" yaml:"-"`
	AlertManager string `json:"alert_manager" yaml:"alert_manager"`
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
	GPS           gps           `json:"gps" yaml:"gps"`
	Mail          mail          `json:"mail" yaml:"mail"`
	RuntimeParam  runtimeParam  `json:"-" yaml:"-"`
}

// Cfg 全局的Config配置，解析dns.yaml的结果
var Cfg *Config

// Version 软件版本
var Version = "1.0.33 bate"

// SubDir for config
var SubDir = "conf.d"
var RuleDir = "rules"

//
var PrometheusConfigName = "prometheus.yml"

//
var configFileName = "config.yml"
