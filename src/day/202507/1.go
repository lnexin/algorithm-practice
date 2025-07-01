package main

import "fmt"

func main() {
	fmt.Println(possibleStringCount("aaaa"))
	//fmt.Println(possibleStringCount("abbcccc"))
}

func possibleStringCount(word string) int {
	rlt := 1
	currentS := ""
	for _, s := range word {
		sc := string(s)
		if currentS == "" {
			currentS = sc
			continue
		} else if currentS == sc {
			rlt++
		} else {
			currentS = sc
		}
	}
	return rlt
}
