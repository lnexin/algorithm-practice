package main

import "fmt"

func main() {
	fmt.Println(findLHS([]int{3999932, 14095060, -38450324, -6006632, -6052261, 34210012, 5636226, -301414, 19348094, 6297539}))
}

func findLHS(nums []int) int {
	maxRlt := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		headNum := nums[i]

		upLen := 0
		downLen := 0
		onlyEqual := true
		for _, n := range nums {
			if n+1 == headNum {
				downLen++
				onlyEqual = false
			} else if n-1 == headNum {
				onlyEqual = false
				upLen++
			} else if n == headNum {
				downLen++
				upLen++
			}
		}

		if onlyEqual {
			continue
		}
		// 存在
		if upLen < downLen {
			upLen = downLen
		}

		if upLen > maxRlt {
			maxRlt = upLen
		}
	}

	return maxRlt

}

// 优化后的findLHS函数
func findLHS2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 使用哈希表统计每个数字的出现次数
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	maxLen := 0

	// 遍历哈希表，检查相邻数字的组合
	for num, count := range countMap {
		// 检查 num+1 是否存在
		if nextCount, exists := countMap[num+1]; exists {
			totalLen := count + nextCount
			if totalLen > maxLen {
				maxLen = totalLen
			}
		}
	}

	return maxLen
}
