package main

import "fmt"

func main() {
	nnn := []int{0, 4, 3, 0, 0}
	fmt.Println(partitionArray(nnn, 0))
}

// 2294. 划分数组使最大差为 K
func partitionArray(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	// 如果要确保有序数组最大差值不超过k
	if len(nums) == 0 {
		return 0
	}

	rlt := 0

	first := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]-nums[first] > k {
			rlt++
			first = i
		}
	}
	if first < len(nums) {
		rlt++
	}
	return rlt
}

func quickSort(nums []int, staInx int, endInx int) {
	if staInx >= endInx {
		return
	}
	// 快速排序
	povit := nums[staInx]

	l := staInx
	r := endInx
	hasChange := false
	for l < r {

		for l < r && nums[r] >= povit {
			// 找到比目标值小的那个下标
			hasChange = true
			r--
		}
		// 循环结束的r 就是那个比目标值小活儿相等的下标
		// 第一次的时候
		nums[l] = nums[r]

		for l < r && nums[l] <= povit {
			hasChange = true
			l++
		}
		// 找到比目标值大的数
		nums[r] = nums[l]
	}

	nums[l] = povit
	if hasChange {
		return
	}
	// 分治
	quickSort(nums, staInx, l-1)
	quickSort(nums, l+1, endInx)
}
