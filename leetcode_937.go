// leetcode 937 重新排列日志文件

package main

import (
	"sort"
	"strings"
	"unicode"
)

// 字母日志排前面，
func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		// Less reports whether the element with index i must sort before the element with index j
		// 简单来说，false就是维持不动，true就是i必须在j前面
		s, t := logs[i], logs[j]
		s1 := strings.SplitN(s, " ", 2)[1] // 以空格分成2部分
		s2 := strings.SplitN(t, " ", 2)[1]
		isDigit1 := unicode.IsDigit(rune(s1[0])) //判断是数字日志还是字符日志
		isDigit2 := unicode.IsDigit(rune(s2[0]))

		if isDigit1 && isDigit2 { //都是数字时，不动
			return false
		}
		if !isDigit1 && !isDigit2 { //都是字符时，排序
			return s1 < s2 || s1 == s2 && s < t
		}
		//如果s1是数字，那么返回false
		//如果s1不是数字，返回true
		return !isDigit1

	})
	return logs
}
