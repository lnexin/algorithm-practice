package main

import "slices"

func main() {

}

func findWordsContaining(words []string, x byte) []int {

	rlt := make([]int, 0)
	for i, word := range words {

		if slices.Contains([]byte(word), x) {
			rlt = append(rlt, i)
		}
	}
	return rlt
}
