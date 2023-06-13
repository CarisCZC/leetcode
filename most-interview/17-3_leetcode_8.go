// https://leetcode.cn/problems/string-to-integer-atoi/
// 8. 字符串转换整数 (atoi)

package main

import "math"

func myAtoi(s string) int {
	numIndex := 0
	flag := true // true 表示正
	flagOK := false
	for i := range s {
		if s[i] == ' ' {
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			numIndex = i
			break
		} else if !flagOK && (s[i] == '-' || s[i] == '+') {
			flagOK = true
			if s[i] == '-' {
				flag = false
			}
			numIndex = i + 1
			break
		} else {
			return 0
		}
	}
	// 此时到达第一个数字
	tmp := 0
	for ; numIndex < len(s); numIndex++ {
		cur := int(s[numIndex] - '0')
		if cur < 0 || cur > 9 {
			break
		}
		tmp = tmp*10 + cur
		if tmp > -math.MinInt32 {
			tmp = -math.MinInt32
			break
		}
	}
	if !flag {
		tmp = -tmp
	}
	if tmp > math.MaxInt32 {
		tmp = math.MaxInt32
	}
	return tmp
}
