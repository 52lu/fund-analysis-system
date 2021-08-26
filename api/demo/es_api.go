package demo

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 索引库
const indexName = "go-test"

// CreateIndex 创建索引
func CreateIndex(ctx *gin.Context) {
	userMapping := `{
    "mappings":{
        "properties":{
            "name":{
                "type":"keyword"
            },
            "age":{
                "type":"byte"
            },
            "birth":{
                "type":"date"
            }
        }
    }
}`
	// 判断索引是否存在
	exist, _ := global.GvaElastic.IndexExists(indexName).Do(ctx)
	if exist {
		response.Error(ctx, "索引已经存在，无需重复创建！")
		return
	}
	res, err := global.GvaElastic.CreateIndex(indexName).BodyString(userMapping).Do(ctx)
	if err != nil {
		response.Error(ctx, "创建索引失败，无需重复创建！")
		return
	}
	response.OkWithData(ctx, res)
}

// SearchById 查询
func SearchById(ctx *gin.Context)  {
	id,_ := ctx.GetQuery("id")
	res, err := global.GvaElastic.Get().Index(indexName).Id(id).Do(ctx)
	if err != nil {
		response.Error(ctx, fmt.Sprintf("查询失败:%s",err))
		return
	}
	response.OkWithData(ctx,res.Source)
}

