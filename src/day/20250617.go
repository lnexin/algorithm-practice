package main

import (
	"strings"
)

// 给你一个字符串 caption，表示一个视频的标题。
//
//需要按照以下步骤 按顺序 生成一个视频的 有效标签 ©leetcode
func generateTag(caption string) string {
	rlt := "#"
	nextUpper := false
	firstLetter := true
	for i := 0; i < len(caption); i++ {
		t := string(caption[i])
		if t == " " {
			if i > 0 {

				nextUpper = true
			}
			continue
		}

		// 转换大小写
		if nextUpper && !firstLetter {
			t = strings.ToUpper(t)

		} else {
			t = strings.ToLower(t)
			firstLetter = false
		}
		nextUpper = false
		rlt += t
	}

	if len(rlt) > 100 {
		rlt = rlt[0:100]
	}
	return rlt
}
