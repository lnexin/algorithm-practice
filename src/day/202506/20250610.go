package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxDifference("mmsmsym"))
}

func maxDifference(s string) int {
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-'a']++
	}

	a1 := 0
	a2 := 0
	for i := 0; i < 26; i++ {
		count := cnt[i]
		if count == 0 {
			continue
		}
		if count%2 == 1 {
			// 奇数 a1
			//a1 = int(math.Max(float64(count), float64(a1)))

			if count > a1 {
				a1 = count
			}
		} else {
			// 偶数个, 去最小的
			if a2 == 0 {
				a2 = count
			}
			if count < a2 {
				a2 = count
			}

			//if count > shortest || shortest == 0 {
			//	shortest = count
			//}
		}
	}

	return a1 - a2
}
