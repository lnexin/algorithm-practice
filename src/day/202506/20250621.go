package main

import "slices"

func main() {

}

func minimumDeletions(word string, k int) int {
	// 字符串中所有数字统计
	freqCnt := make([]int, 26)
	for i := range word {
		freqCnt[i-'a']++
	}
	// 出现频率之差与k的关系

	cnt := make([]int, 0)
	// <=k
	for _, n := range freqCnt {
		if n > 0 {
			cnt = append(cnt, n)
		}
	}
	slices.Sort(cnt)

	if cnt[len(cnt)-1]-cnt[0] <= k {
		return 0
	}

	// 如果需要删除
	rlt := len(word)
	for _, a := range cnt {
		// 选定一个字母
		steps := 0
		for _, b := range cnt {
			if a > b {
				steps += b
			} else if a+k < b {
				steps += b - (a - k)
			}
		}
		if steps < rlt {
			rlt = steps
		}
	}
	return rlt
}

func pass(cnts []int, k int) bool {
	cnt := make([]int, 0)
	// <=k
	for _, n := range cnts {
		if n > 0 {
			cnt = append(cnt, n)
		}
	}
	slices.Sort(cnt)

	return cnt[len(cnt)-1]-cnt[0] <= k
}
