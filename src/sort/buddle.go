package main

import "fmt"

// 冒泡排序

func main() {
	fmt.Println(sortBuddle([]int64{16, 5, 1, 4, 1}))
}
func sortBuddle(nums []int64) []int64 {
	length := len(nums)
	if length <= 1 {
		return nums
	}

	head := 0
	for head < length-1 {
		changed := false
		tail := length - 1
		for tail > head {
			if nums[tail] < nums[tail-1] {
				temp := nums[tail-1]
				nums[tail-1] = nums[tail]
				nums[tail] = temp

				changed = true
			}
			tail--
		}
		if !changed {
			break
		}
		head++
	}
	return nums
}
