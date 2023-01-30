// https://leetcode.cn/problems/longest-palindromic-substring/?favorite=2ckc81c
package main

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= i; j-- {
			if IsPalindrome(s[i : j+1]) {
				if len(s[i:j+1]) > len(res) {
					res = s[i : j+1]
				}
			}
		}
	}
	return res
}

func IsPalindrome(s string) bool {
	for i, j := 0, len(s)-1; j >= i; {
		if s[j] == s[i] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
