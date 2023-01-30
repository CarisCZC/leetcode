// https://leetcode.cn/problems/string-to-integer-atoi/?favorite=2ckc81c
package main

import "math"

func myAtoi(s string) int {
	maxNum := math.MaxInt32
	minNum := math.MinInt32
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if (s[i] == '-' || s[i] == '+') && i < len(s)-1 && isNum(s[i+1]) {
			continue
		}
		if isNum(s[i]) {
			for j := i; j < len(s); j++ {
				if res > maxNum {
					break
				}
				if isNum(s[j]) {
					res *= 10
					res += int(s[j]) - 48
				} else {
					break
				}
			}
			// 判断正负
			if i > 0 && s[i-1] == '-' {
				res = 0 - res
			}
			break
		} else {
			break
		}

	}
	if res > maxNum {
		return maxNum
	}
	if res < minNum {
		return minNum
	}
	return res
}

func isNum(ch byte) bool {
	if ch >= 48 && ch <= 57 {
		return true
	} else {
		return false
	}
}

func main() {
	myAtoi("922337299999")
}
