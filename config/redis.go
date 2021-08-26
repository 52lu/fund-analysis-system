/**
 * @Description redis配置
 **/
package config

import "time"

type redis struct {
	Addr        string        `yaml:"addr"`
	Password    string        `yaml:"password"`
	DefaultDB   int           `yaml:"defaultDB"`
	DialTimeout time.Duration `yaml:"dialTimeout"`
	Enable      bool          `yaml:"enable"`
}
