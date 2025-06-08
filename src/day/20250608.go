package main

import (
	"fmt"
)

func main() {
	fmt.Println(lexicalOrder(5656))
}

func lexicalOrder(n int) []int {
	rlt := make([]int, 0)

	num := 1
	for len(rlt) < n {
		rlt = append(rlt, num)
		// 评断下一个数字是多少

		// 如果上一个数字有效, 那么末尾 + 0, 就是最小的
		if num*10 <= n {
			num = num * 10
			continue
		}

		// 如果不行, 那理论上是要 + 1
		// 当时如果这个数字大于n, 那么这个数字就无效了, 需要去处末尾一位 然后 +1
		//if num+1 > n {
		//	num /= 10
		//}
		// 同时如果末尾为9, 那后面+1 的话, 变为0之后, 那会缺少此时无零的字典序
		// 199 -> 200, 下一个应该是 2, 而不是20
		// 119 -> 120, 下一个应该是 12 不是 120
		// 1199 -> 1200, 下一个应该是 12, 不是 120
		// 所以使用循环, 去掉末尾9导致进位
		//for num%10 == 9 {
		//	num /= 10
		//}

		// 整合写法
		for num%10 == 9 || num > n {
			num /= 10
		}
		num++
	}
	return rlt
}

//func lexicalOrder(n int) []int {
//	if n <= 1 {
//		return []int{1}
//	}
//
//	rlt := make([]int, 0)
//
//	i := 1
//	for len(rlt) < n {
//		//fmt.Println(n)
//		if i == 1099 {
//			fmt.Println(i)
//		}
//
//		// 如果此时 新数字大于n, 就代表这个数字不在范围内, 首位进一, 开始处理
//		oldFirst := getIndexNum(i, 0)
//
//		if i > n {
//			i = oldFirst + 1
//			if i == 10 {
//				break
//			}
//			continue
//		} else {
//			rlt = append(rlt, i)
//		}
//		// 这里面有个逻辑
//		// 10 -> 11 不对 10 -> 100
//		if i*10 <= n {
//			i = i * 10
//			continue
//		}
//
//		// 更新i
//		newFirst := getIndexNum(i+1, 0)
//
//		// 首位若相等, + 1 进行下次处理, 若下次的 i 大于n, 就代表是真的结束了
//		if oldFirst == newFirst {
//
//			// 109 -> 110
//			// 这种进位的情况 去 0
//			if getIndexNum(i, -1) == 9 && getIndexNum(i+1, -1) == 0 {
//
//				//i = (i + 1) / 10
//				i = i + 1
//				for i%10 == 0 {
//					i = i / 10
//				}
//
//				continue
//			}
//			if i+1 > n {
//				// 这个时候也要进位
//				i = i/10 + 1
//				//
//				if getIndexNum(i, 0) != oldFirst {
//					i = getIndexNum(i, 0)
//				}
//
//				continue
//			}
//			i++
//			continue
//		}
//
//		// 此时首位不相等
//		// 1+1 =2不对 -> 1 -> 10
//
//		// 如果 i + 1之后, 第一位为不相等, 代表进位了 199->200, 此时需要重置为 1000
//		if oldFirst != newFirst {
//			l := len(strconv.Itoa(i))
//			i = oldFirst * int(math.Pow10(l))
//			if slices.Contains(rlt, i) {
//				i = newFirst
//				continue
//			}
//		}
//		// 如果 i + 1之后, 第一位为1, 进行下一次循环 (此时仍然小于n)
//		// 如果 i+1 之后, 第一位不为1 就保留首位1, 右移一位, (此时仍然小于n)                                                ）
//
//		// 都不满足的时候, 首位 +1 左移 n-1位 (此时仍然小于n)
//
//		// 都不满足的时候, 首位 +1 左移 n-1位 (此时仍然小于n)
//	}
//	return rlt
//}
//
//func getIndexNum(i int, index int) int {
//	if i < 10 {
//		return i
//	}
//
//	// 获取首位数字
//	s := strconv.Itoa(i)
//	if index >= len(s) {
//		index = len(s) - 1
//	}
//	if index == -1 {
//		index = len(s) - 1
//	}
//	c := s[index]
//	atoi, _ := strconv.Atoi(string(c))
//	return atoi
//}
