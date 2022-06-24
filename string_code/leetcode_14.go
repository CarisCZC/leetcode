package main

// 1. 选数组第一个元素为flag
// 2. 其他字符串，比较时选最短的为n
func longestCommonPrefix(strs []string) string {
	flag := strs[0]
	n := len(flag)
	minlen := len(flag)
	for i := 1; i < len(strs); i++ {
		if minlen > len(strs[i]) {
			minlen = len(strs[i])
		}
		j := 0
		for j < minlen {
			if flag[j] == strs[i][j] {
				j++
			} else {
				break
			}
		}
		if j < n {
			n = j
		}
	}
	return flag[:n]
}
