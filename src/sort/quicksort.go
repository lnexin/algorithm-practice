package main

func main() {

}

func quickSort(nums []int, staInx int, endInx int) {
	// 快速排序
	povit := nums[0]

	l := 0
	r := len(nums) - 1
	for l < r {
		for l < r && nums[r] >= povit {
			// 找到比目标值小的那个下标
			r--
		}
		// 循环结束的r 就是那个比目标值小活儿相等的下标
		// 第一次的时候
		nums[l] = nums[r]

		for l < r && nums[l] <= povit {
			l++
		}
		// 找到比目标值大的数
		nums[r] = nums[l]
	}

	nums[l] = povit
	// 分治
	quickSort(nums, staInx, l-1)
	quickSort(nums, l+1, endInx)
}
