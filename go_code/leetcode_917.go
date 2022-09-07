// leetcode 917 ,仅仅反转字母

package main

import "unicode"

// 前后双指针，遇到非字母就跳过
func reverseOnlyLetters(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		for i < j && !unicode.IsLetter(runes[i]) {
			i++
		}
		for i < j && !unicode.IsLetter(runes[j]) {
			j--
		}
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
