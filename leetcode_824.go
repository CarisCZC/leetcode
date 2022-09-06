// leetcode 824, 山羊拉丁文
package main

import "strings"

//1. 取单词
//2. 做修改
//3. 添加
func toGoatLatin(sentence string) string {
	pre := 0
	flag := true
	res := strings.Builder{}
	cnt := 1
	yuan := []byte{'a', 'e', 'i', 'o', 'u','A','E','I','O','U'}
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' && flag {
			isyuan := false
			for _, v := range yuan {
				if sentence[pre] == v {
					isyuan = true
					break
				}
			}
			if isyuan {
				res.WriteString(sentence[pre:i] + "ma")
			} else {
				res.WriteString(sentence[pre+1:i] + sentence[pre:pre+1] + "ma")
			}
			for j := 0; j < cnt; j++ {
				res.WriteString("a")
			}
			cnt++
			res.WriteString(" ")
			flag = false
		}
		if sentence[i] != ' ' && !flag {
			pre = i
			flag = true
		}
		if i == len(sentence)-1 {
			isyuan := false
			for _, v := range yuan {
				if sentence[pre] == v {
					isyuan = true
					break
				}
			}
			if isyuan {
				res.WriteString(sentence[pre:i+1] + "ma")
			} else {
				res.WriteString(sentence[pre+1:i+1] + sentence[pre:pre+1] + "ma")
			}
			for j := 0; j < cnt; j++ {
				res.WriteString("a")
			}
		}
	}
	return res.String()
}
