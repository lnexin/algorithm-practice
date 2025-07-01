package main

import "slices"

func main() {

}
func numSubseq(nums []int, target int) int {
	l := len(nums)
	p := int(1e9) + 7
	f := make([]int, l)
	f[0] = 1
	for i := 1; i < l; i++ {
		f[i] = f[i-1] * 2 % p
	}

	slices.Sort(nums)

	left, right := 0, l-1
	res := 0
	for left <= right {
		if nums[left]+nums[right] <= target {
			res += f[right-left]
			left++
		} else {
			right--
		}

	}
	//for left <= right {
	//	if nums[left] + nums[right] > target {
	//		right--
	//	} else {
	//		res = (res + f[right-left]) % p
	//		left++
	//	}
	//}
	return res % p

}
