// leetcode 804,唯一摩尔斯密码词
package main

import "strings"

func uniqueMorseRepresentations(words []string) int {
	ch := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..",
		".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-",
		"..-", "...-", ".--", "-..-", "-.--", "--.."}
	tmp := map[string]bool{}
	for _, v := range words {
		mosi := strings.Builder{}
		for _, k := range v {
			mosi.WriteString(ch[k-'a'])
		}
		tmp[mosi.String()] = true
	}
	return len(tmp)
}
