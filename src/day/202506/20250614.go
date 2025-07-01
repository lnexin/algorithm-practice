package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(minMaxDifference(99999))
}

func minMaxDifference(num int) int {
	// 最大, 首位替换成9
	// 最小, 首位替换成 0
	textNum := strconv.Itoa(num)

	firstNum := textNum[0]
	// 要找最大的, 需要找到从左往右第一位不是9的
	firstMaxNum := firstNum
	i := 1
	for firstMaxNum == '9' && i < len(textNum) {
		firstMaxNum = textNum[i]
		i++
	}
	fmt.Println(textNum, firstMaxNum, firstNum)

	mm := strings.ReplaceAll(textNum, string(firstMaxNum), "9")
	mi := strings.ReplaceAll(textNum, string(firstNum), "0")

	//mm := ""
	//mi := ""
	//for i := 0; i < len(textNum); i++ {
	//	cur := textNum[i]
	//	if cur != firstNum && cur != firstMaxNum {
	//		mm += string(textNum[i])
	//		mi += string(textNum[i])
	//		continue
	//	}
	//
	//	if cur == firstMaxNum {
	//		mm += "9"
	//	}
	//	if cur == firstNum {
	//		mi += "0"
	//	}
	//}

	mmI, err := strconv.Atoi(mm)
	if err != nil {
		return 0
	}
	miI, err := strconv.Atoi(mi)
	fmt.Println(mmI, miI)
	return mmI - miI
}
