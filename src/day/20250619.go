package main

import (
	"fmt"
	"slices"
)

func main() {
	nnn := []int{0, 4, 3, 0, 0}
	fmt.Println(partitionArray(nnn, 0))
}

// 2294. 划分数组使最大差为 K
func partitionArray(nums []int, k int) int {
	//quickSort(nums, 0, len(nums)-1)
	slices.Sort(nums)
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

// 3路处理只能减少重复元素的问题
func quickSort2(nums []int, left int, right int) []int {
	if len(nums) <= 1 {
		return nums
	}

	l := left
	p := nums[l]
	// 分 3组, 大于, 等于, 小于
	gatherArr := []int{}
	equalArr := []int{}
	lessArr := []int{}
	for i := range nums {
		if nums[i] > p {
			gatherArr = append(gatherArr, nums[i])
		} else if nums[i] == p {
			equalArr = append(equalArr, nums[i])
		} else {
			lessArr = append(lessArr, nums[i])
		}
	}

	lessRlt := quickSort2(lessArr, left, len(lessArr)-1)
	lessRlt = append(lessRlt, equalArr...)
	lessRlt = append(lessRlt, quickSort2(gatherArr, 0, len(gatherArr)-1)...)
	return lessRlt
}
