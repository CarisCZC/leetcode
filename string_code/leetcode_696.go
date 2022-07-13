// leetcode 696, 计数子字符串

// 字符串有相同数量的0和1，且0，1都是连续的
package main

// 先统计连续的0，1的坐标
func countBinarySubstrings(s string) int {
	counts := []int{}
	counts = append(counts, 1)
	n := len(counts)
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			counts[n-1]++
		} else {
			counts = append(counts, 1)
			n++
		}
	}
	res := 0
	for i := 0; i < n-1; i++ {
		res += min(counts[i], counts[i+1])
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
