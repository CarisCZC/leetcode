// leetcode 482,密钥格式化

package main

import "strings"

func licenseKeyFormatting(s string, k int) string {
	flag := 0
	res := ""
	for i := len(s); i > 0; i-- {
		if s[i-1] != '-' {
			if flag == k {
				res = "-" + res
				flag = 0
			}
			if flag < k {
				res = strings.ToTitle(string(s[i-1])) + res
				flag++
			}
		}
	}
	return res
}

func main() {
	s := "2-5g-3-J"
	licenseKeyFormatting(s, 2)
}
