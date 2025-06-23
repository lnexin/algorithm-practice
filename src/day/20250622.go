package main

import "fmt"

func main() {
	fmt.Println(divideString("abcdefghij", 3, 'x'))

}

func divideString(s string, k int, fill byte) []string {
	rlt := make([]string, 0)

	for i := 0; i < len(s); i += k {

		end := i + k
		if end < len(s) {
			rlt = append(rlt, s[i:end])
		} else {
			// >= 要补 end -len个 fill
			text := s[i:len(s)]
			for j := len(text); j < k; j++ {
				text += string(fill)
			}
			rlt = append(rlt, text)
		}
	}
	return rlt
}
