package test

import (
	"fmt"
	"testing"
)

// 快排
func quickSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	benchVal := data[0]
	var rightArr, leftArr []int
	// 从第二个数开始遍历
	for i := 1; i < len(data); i++ {
		if data[i] < benchVal {
			leftArr = append(leftArr, data[i])
		} else {
			rightArr = append(rightArr, data[i])
		}
	}
	// 分组后继续操作
	rightArr = quickSort(rightArr)
	leftArr = quickSort(leftArr)

	// 合并结果
	leftArr = append(leftArr, benchVal)
	for i := 0; i < len(rightArr); i++ {
		leftArr = append(leftArr, rightArr[i])
	}
	return leftArr
}

func maoSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data); j++ {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
	return data
}

func TestRu(t *testing.T) {
	d := []int{3, 15, 9, 12, 1, 2, 55, 29}
	dd := quickSort(d)
	fmt.Println("排序结果：",dd)
	i := search2(12, dd)
	fmt.Println("index = ", i)
	//fmt.Println(quickSort(d))
	//fmt.Println(maoSort(d))
}

// [ m ]
func search2(search int, data []int) int {
	beginIndex := 0
	endIndex := len(data) - 1
	for beginIndex <= endIndex {
		middleIndex := (beginIndex + endIndex) / 2
		if search == data[middleIndex] {
			return middleIndex
		} else if search < data[middleIndex] {
			endIndex = middleIndex - 1
		} else {
			beginIndex = middleIndex + 1
		}
	}
	return  -1
}
