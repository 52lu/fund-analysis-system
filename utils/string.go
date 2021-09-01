package utils

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"regexp"
)

// 提取字符串中的数字
func ExtractNumberFromString(str string) string {
	// 定义正则匹配规则
	var compile = regexp.MustCompile(`\d+\.?\d+`)
	match := compile.FindAllStringSubmatch(str, -1)
	numberString := ""
	for _, val := range match {
		if len(val) > 0 {
			numberString = numberString + val[0]
		}
	}
	return numberString
}

// GbkToUtf8 GBK转成UTF8
func GbkToUtf8(str string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), simplifiedchinese.GBK.NewDecoder())
	all, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(all), nil
}

