// test: 临时测试
package test

import (
	_ "52lu/fund-analye-system"
	"52lu/fund-analye-system/global"
	"fmt"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	fmt.Printf("%+v\n",global.GvaConfig)
	str := "Phone格式不正确！"
	fmt.Println("str = ",strings.ToLower(str))
}
