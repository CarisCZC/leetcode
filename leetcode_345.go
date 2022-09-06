// leetcode 345,反转字符中的元音字母

// 依旧是双指针，加入判断即可
package main

import "strings"

func reverseVowels(s string) string {
	vowels := "aeiou"
	bytes := []byte(s)
	i := 0
	j := len(s) - 1
	for i < j {
		for i < j && !strings.Contains(vowels, string(s[i])) {
			i++
		}
		for i < j && !strings.Contains(vowels, string(s[j])) {
			j--
		}
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--

	}
	return string(bytes)
}

func main() {
	s := "hello"
	reverseVowels(s)
}
