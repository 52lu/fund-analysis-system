/**
 * @Description elasticSearch配置
 **/
package config

import "time"

/*
elastic:
  # 是否开启
  enable: true
  # 服务地址，多个地址用逗号隔开
  url: http://127.0.0.1:9200
  # 是否转换请求地址，默认为true,当等于true时 请求http://ip:port/_nodes/http，将其返回的url作为请求路径
  sniff: false
  # 心跳检测间隔
  healthCheckInterval: 5s
  # 日志前缀
  logPre: ES-
*/
type elastic struct {
	Url                 string        `yaml:"url"`
	Sniff               bool          `yaml:"sniff"`
	HealthCheckInterval time.Duration `yaml:"healthCheckInterval"`
	LogPre              string        `yaml:"logPre"`
	Enable              bool          `yaml:"enable"`
}
