package main

func main() {
	num1 := []int{-2, -1, 0, 1, 2}
	num2 := []int{-3, -1, 2, 4, 5}
	kthSmallestProduct(num1, num2, 3)
}

func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	len1 := int64(len(nums1))
	len2 := int64(len(nums2))

	// 第 k 小
	iStart := k / len2
	if iStart < 1 {
		iStart = 1
	}
	mod := k % len2

	inx1 := iStart - 1
	inx2 := mod - 1

	minRlt := int64(0)
	// 分别用两个序列去两盒nums里面取
	if inx1 < len1 && inx2 < len2 {
		minRlt = int64(nums1[inx1] * nums2[inx2])
	}

	if inx1 < len2 && inx2 < len1 {
		if newRlt := int64(nums1[inx2] * nums2[inx1]); newRlt < minRlt {
			minRlt = newRlt
		}
	}

	//for i := iStart - 1; i < len1; i++ {
	//	for j := mod - 1; j < len2; j++ {
	//		if (i+1)*(j+1) == k {
	//			rlt := nums1[i] * nums2[j]
	//			return int64(rlt)
	//		}
	//	}
	//
	//}
	return minRlt
}
