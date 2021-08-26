/**
 * @Description elasticSearch 客户端创建
 **/
package initialize

import (
	"52lu/fund-analye-system/global"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

// 创建es客户端
func initElastic() {
	// 配置
	elasticConfig := global.GvaConfig.Elastic
	if elasticConfig.Enable {
		fmt.Printf("elasticConfig: %v\n", elasticConfig)
		// 创建客户端
		client, err := elastic.NewClient(
			elastic.SetURL(elasticConfig.Url),
			elastic.SetSniff(elasticConfig.Sniff),
			elastic.SetHealthcheckInterval(elasticConfig.HealthCheckInterval),
			elastic.SetErrorLog(log.New(os.Stderr, elasticConfig.LogPre+"ERROR ", log.LstdFlags)),
			elastic.SetInfoLog(log.New(os.Stderr, elasticConfig.LogPre+"INFO ", log.LstdFlags)),
		)
		if err != nil {
			panic("创建ES客户端错误:" + err.Error())
		}
		global.GvaElastic = client
	}
}
