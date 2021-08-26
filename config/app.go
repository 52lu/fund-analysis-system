// Package config /**
package config

// 应用信息
type app struct {
	Desc       string `yaml:"desc"`
	Addr       string `yaml:"addr"`
	ConfigFile string `yaml:"configFile"`
	Version    string `yaml:"version"`
	Env        string `yaml:"env"`
}

// ServerConfig 配置信息
type ServerConfig struct {
	App     app     `yaml:"app"`
	Mysql   mysql   `yaml:"mysql"`
	Log     log     `yaml:"log"`
	Jwt     jwt     `yaml:"jwt"`
	Redis   redis   `yaml:"redis"`
	Elastic elastic `yaml:"elastic"`
	Cron    cron    `yaml:"cron"`
}
