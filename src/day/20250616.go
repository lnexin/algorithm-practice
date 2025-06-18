package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumDifference([]int{999, 997, 980, 976, 948, 940, 938, 928, 924, 917, 907, 907, 881, 878, 864, 862, 859, 857, 848, 840, 824, 824, 824, 805, 802, 798, 788, 777, 775, 766, 755, 748, 735, 732, 727, 705, 700, 697, 693, 679, 676, 644, 634, 624, 599, 596, 588, 583, 562, 558, 553, 539, 537, 536, 509, 491, 485, 483, 454, 449, 438, 425, 403, 368, 345, 327, 287, 285, 270, 263, 255, 248, 235, 234, 224, 221, 201, 189, 187, 183, 179, 168, 155, 153, 150, 144, 107, 102, 102, 87, 80, 57, 55, 49, 48, 45, 26, 26, 23, 15}))
}

// 计算 nums[j] - nums[i] 能求得的 最大差值 ，其中 0 <= i < j < n 且 nums[i] < nums[j] 。
// 返回 最大差值 。如果不存在满足要求的 i 和 j ，返回 -1
func maximumDifference(nums []int) int {
	i := 0
	j := 1
	l := len(nums)

	if l < 2 {
		return -1
	}
	maxRlt := -1
	for j < l || l > i+1 {

		if j >= l {
			i++
			j = i + 1
		}
		if i+1 >= l {
			break
		}
		//fmt.Println(i, nums[i], j, nums[j])
		diff := nums[j] - nums[i]

		// 两个数
		// i + 1 的条件
		if diff > 0 && diff > maxRlt {
			fmt.Println(i, j, nums[j])
			fmt.Println(nums[j], nums[j], diff)
			maxRlt = diff
		}
		j++
	}

	return maxRlt
}
