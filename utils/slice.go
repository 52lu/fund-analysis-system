package utils

// 判断字符串是否在切片中
func ExistSliceStr(search string, sliceStr []string) (int, bool) {
	for k, v := range sliceStr {
		if search == v {
			// 找到
			return k, true
		}
	}
	return -1, false
}

