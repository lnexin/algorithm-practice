package main

import "math"

func main() {

}
func maxSubsequence(nums []int, k int) []int {
	// 结果
	rlt := make([]int, 0)
	curMin := math.MaxInt
	curMinInx := 0
	//curMax := math.MinInt
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		if len(rlt) < k {
			rlt = append(rlt, cur)
			//if cur > curMax {
			//	curMax =cur
			//}
			if cur < curMin {
				curMin = cur
				curMinInx = len(rlt) - 1
			}

			continue
		}
		// 如果长度已经满足了
		// 需要比较有没有比他小的
		if cur > curMin {
			// 如果当前值比序列中最小的数大,替换
			
		}
	}

	// k个
	for i := 0; i < k; i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}

	// 取前k个
	subNum := nums[:k]
	rlt := 0
	for _, n := range subNum {
		rlt += n
	}
	return rlt
}
