package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-5, -10, -5}
	fmt.Println(maxAdjacentDistance(nums))
}

// 3423. 循环数组中相邻元素的最大差值
func maxAdjacentDistance(nums []int) int {
	//
	rlt := 0
	// 相邻
	for i := range nums {
		a1 := nums[i]
		a2 := a1
		if i == len(nums)-1 {
			a2 = nums[0]
		} else {
			a2 = nums[i+1]
		}
		sub := int(math.Abs(float64(a1 - a2)))

		fmt.Println(a1, a2, sub, rlt)
		if rlt < sub {
			rlt = sub
		}
	}
	return rlt
}
