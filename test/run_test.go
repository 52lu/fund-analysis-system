// test: 临时测试
package test

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/initialize"
	"fmt"
	"strings"
	"testing"
)


func TestRun(t *testing.T) {
	initialize.SetLoadInit()
	fmt.Printf("%+v\n",global.GvaConfig)
	str := "Phone格式不正确！"
	fmt.Println("str = ",strings.ToLower(str))
}
