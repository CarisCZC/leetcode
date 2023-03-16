package main

// KMP算法(不会)
// 直接截取，也不失为一个好方法
func strStr(haystack string, needle string) int {
	n := len(needle)
next:
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] && haystack[i+n-1] == needle[n-1] {
			for j := 1; j <= n-2; j++ {
				if needle[j] != haystack[i+j] {
					continue next
				}
			}
			return i
		}
	}
	return -1
}
