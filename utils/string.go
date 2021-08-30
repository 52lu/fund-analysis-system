package utils

import "regexp"

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
