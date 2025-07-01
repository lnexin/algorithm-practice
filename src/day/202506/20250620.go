package main

import "math"

func main() {

}
func maxDistance(s string, k int) int {
	length := len(s)
	x := 0
	y := 0
	rlt := 0

	for i := 0; i < length; i++ {
		switch s[i] {
		case 'N':
			y++
		case 'S':
			y--
		case 'W':
			x--
		case 'E':
			x++
		}
		abs := int(math.Abs(float64(x)) + math.Abs(float64(y)))
		// 当前的距离
		rlt = int(math.Max(float64(rlt), float64(abs+2*k)))
	}

	// 此时的x和y的绝对值已经是距离了
	// 若k步, 那么给
	return rlt

}
