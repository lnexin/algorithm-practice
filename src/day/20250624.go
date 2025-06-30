package main

func main() {
	findKDistantIndices([]int{3, 4, 9, 1, 3, 9, 5}, 9, 1)
}

func findKDistantIndices(nums []int, key int, k int) []int {
	rlt := make([]int, len(nums))
	// 遍历找到与 k 相等的数字
	for j, num := range nums {
		if num == key {
			// 找到 i - j <= k 的下标i
			//for i := range nums {
			//	abs := math.Abs(float64(i - j))
			//	if int(abs) <= k {
			//		rlt[i] = 1
			//	}
			//}
			start := j - k
			if start < 0 {
				start = 0
			}
			end := j + k
			if end >= len(nums) {
				end = len(nums) - 1
			}
			for i := start; i <= end; i++ {
				rlt[i] = 1
			}
		}
	}

	r := make([]int, 0)
	for i := range rlt {
		if rlt[i] > 0 {
			r = append(r, i)
		}
	}
	return r
}
