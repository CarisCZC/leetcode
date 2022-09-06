// leetcode 680,验证回文字符串II

// 给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串

package main

func validPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	flag1, flag2 := true, true
	// cnt := 1 // 有一个删除的额度
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			for k, m := i+1, j; flag1 && k < m; k, m = k+1, m-1 {
				if s[k] != s[m] {
					flag1 = !flag1
				}
			}
			if !flag1 {
				for k, m := i, j-1; flag2 && k < m; k, m = k+1, m-1 {
					if s[k] != s[m] {
						flag2 = !flag2
					}
				}
			}
			return flag1 || flag2
		}
	}
	return true
}

func main() {
	s := "abc"
	validPalindrome(s)
}
