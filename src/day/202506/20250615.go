package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(maxDiff(1101057))
}

func maxDiff(num int) int {

	//
	text := strconv.Itoa(num)
	length := len(text)
	// 找最大的, 第一位替换为9, 若第一位为9, 找到第一位不为九的数字
	maxFirst := text[0]
	i := 0
	for maxFirst == '9' && i < length {
		maxFirst = text[i]
		i++
	}

	// 最小值的时候, 不能处理第一位为 0
	// 所以如果结构通第一位相同, 那么只能处理为1
	minFirst := text[0]
	// 如果第一位为1, 那么就要往后找不为1的那一位
	for (minFirst == '0' || minFirst == '1') && i < length {
		minFirst = text[i]
		i++
	}
	targetMin := "0"
	if minFirst == text[0] {
		targetMin = "1"
	}

	// 如果最大值的时候, 首位替换成9
	// 如果最小值的时候, 首位替换成0
	maxRlt := strings.ReplaceAll(text, string(maxFirst), "9")
	minRlt := strings.ReplaceAll(text, string(minFirst), targetMin)
	maxInt, err := strconv.Atoi(maxRlt)
	if err != nil {
		return 0
	}
	minInt, err := strconv.Atoi(minRlt)
	if err != nil {
		return 0
	}
	diff := maxInt - minInt
	return diff

}
