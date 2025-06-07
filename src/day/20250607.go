package main

import (
	"fmt"
)

func main() {
	fmt.Println(clearStars("d*o*"))
	//fmt.Println(clearStars("aada*b"))
}

// 20250607

func clearStars(s string) string {
	lenght := len(s)
	charLastIndex := make([]int, 26)
	for i := range charLastIndex {
		charLastIndex[i] = -1
	}

	sArr := []rune(s)

	// 每个字符前一个相同的字符下表
	pervSameCharInx := make([]int, lenght)
	minChar := 'z' + 1
	//minCharInx := math.MaxInt
	for i, c := range sArr {

		if c != '*' {
			cInx := c - 'a'
			pervSameCharInx[i] = charLastIndex[cInx]
			// 赋值以后就更新
			charLastIndex[cInx] = i
			if c < minChar {
				minChar = c
				//fmt.Println("update minchar:", minChar, string(minChar))
			}
		} else {
			// 星号删除
			sArr[i] = 0
			// 删除左侧最小字符的最后一位
			minCharLastInx := minChar - 'a'
			//fmt.Println(minCharLastInx, string(minChar))
			sArr[charLastIndex[minCharLastInx]] = 0
			// 更新最小字符
			charLastIndex[minCharLastInx] = pervSameCharInx[charLastIndex[minCharLastInx]]

			// 如果 minchar 在左侧只有一个, 经过上面的删除, SArr[minCharLastInx] 会变成0, charLastIndex[minCharLastInx] 会变成-1
			// 那现在这个minChar 就是无效的, 需要重新寻找
			// 这个时候就要在 lastCharInx 里面找到一个最小的
			if charLastIndex[minCharLastInx] == -1 {
				minChar = 'z' + 1

				for i2, v := range charLastIndex {
					// 如果出现不为1 的最小就更新
					if v != -1 && minChar > 'a'+rune(i2) {
						minChar = 'a' + rune(i2)
					}
				}
			}
		}

	}

	rlt := ""
	for i := 0; i < lenght; i++ {
		if sArr[i] != 0 {
			rlt += string(sArr[i])
		}
	}

	return rlt
}

// 1. 删除所有星号
// =================================================================
//func clearStars(s string) string {
//	cnt := make([]int64, 26)
//
//REDO:
//	leftMinChar := 'z' + 1
//	leftMinCharInx := -1
//	for i, c := range s {
//		fmt.Println(i, string(c), c)
//		if c != '*' && c <= leftMinChar {
//			leftMinChar = c
//			leftMinCharInx = i
//		}
//		if c == '*' {
//			lastChar := ""
//			if i < len(s)-1 {
//				lastChar = s[i+1:]
//			} else {
//
//				// 最后一个是星号
//			}
//			s = s[:leftMinCharInx] + s[leftMinCharInx+1:i] + lastChar
//			// 先删除当前的星号
//			//s = s[:i] + s[i+1:]
//			// 删除左侧最小的字符, 因为这个字符肯定在星号左边
//			//s = s[:leftMinCharInx] + s[leftMinCharInx+1:]
//			goto REDO
//		}
//	}
//	// 输出最小字典序
//	rlt := ""
//	for i := 0; i < len(cnt); i++ {
//		if cnt[i] == 0 {
//			continue
//		}
//		s += string('a' + i)
//		for j := int64(0); j < cnt[i]; j++ {
//			rlt = rlt + s
//		}
//	}
//	return s
//
//}

// =================================================
//func clearStars(s string) string {
//
//	// while is * , do
//	// delete min char in left sub-char
//
//	cnt := make([]int64, 26)
//	starCnt := 0
//	lastStarIndex := -1
//
//	for i, c := range s {
//		if c == '*' {
//			starCnt++
//			if i >= lastStarIndex {
//				lastStarIndex = i
//			}
//		} else {
//			cnt[c-'a']++
//		}
//	}
//
//	if starCnt == 0 {
//		return s
//	}
//	// 删除最后一个星号左边的最小字符
//	// 要看最小字符是不是在星号左边
//
//	leftMinChar := uint8('z' + 1)
//
//	for i := lastStarIndex - 1; i >= 0; i-- {
//		if leftMinChar == 'a' {
//			break
//		}
//		c := s[i]
//		if leftMinChar > c {
//			leftMinChar = c
//			//fmt.Println(string(leftMinChar))
//		}
//	}
//
//	result := ""
//	// 输出news的最小字典序, 其字符串的个数在cnt已有统级
//	for i := 0; i < len(cnt); i++ {
//		if cnt[i] == 0 {
//			continue
//		}
//		// cnt[0] 是a的个数, cnt[1]是b的个数, 依次类推
//		c := uint8('a' + i)
//		fmt.Println(string(c), cnt[i])
//		if c == leftMinChar {
//			cnt[i]--
//		}
//		fmt.Println(string(c), cnt[i])
//		// 添加cnt[i]个c
//		for j := int64(0); j < cnt[i]; j++ {
//			result += string(c)
//		}
//
//	}
//
//	return result
//}
