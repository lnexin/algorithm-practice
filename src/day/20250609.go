package main

import "fmt"

func main() {
	fmt.Println(findKthNumber(13, 2))
}

func findKthNumber(n int, k int) int {
	// 1 2 3 4 5 6 7 8 9 10

	curNum := 1
	//k--
	for k > 0 {
		curSteps := calcSteps(curNum, n)
		// 当前节点下的数据
		// 如果当前节点下的数据数量小于k,那么k个一定不在这个节点里面
		// 反之在这个节点里面
		if curSteps < k {
			k -= curSteps
			curNum++
		} else if curSteps == k {
			return curNum
		} else {
			k--
			curNum *= 10
			// 在这个节点里面
		}
	}
	return curNum

}

func calcSteps(num int, n int) int {
	// 给一个节点, 计算其下有多少个节点

	// 不能大于 n
	// 每一个点下面 10 -> 11,12,13,14,15,16,17,18,19

	steps := 0

	first := num
	last := num

	for first <= n {
		if last > n {
			last = n
		}
		steps += last - first + 1
		first = first * 10
		last = last*10 + 9

	}
	return steps
}
