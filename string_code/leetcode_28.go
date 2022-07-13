// leetcode 28 , 实现strStr()
// 即在字符串中找子串
package main

// haystack中找到needle的第一个字母出现的位置
// 然后截取needle长度，和needle作比较
func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 || n > len(haystack) {
		return 0
	}
	for i := 0; i < len(haystack)-n; i++ {
		if haystack[i] == needle[0] {
			tmp := haystack[i : i+n]
			if tmp == needle {
				return i
			}
		}
	}
	return -1
}
