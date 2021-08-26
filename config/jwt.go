/**
 * @Description JWT配置
 **/
package config

import "time"

// JSON WEB TOKEN 配置
type jwt struct {
	Secret string        `yaml:"secret"`
	Issuer string        `yaml:"issuer"`
	Expire time.Duration `yaml:"expire"`
}
