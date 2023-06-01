// https://leetcode.cn/problems/longest-common-prefix/
// 14. 最长公共前缀

package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	idx := 0
	var cur byte
	for {
		for i := 0; i < len(strs); i++ {
			if idx >= len(strs[i]) {
				return strs[0][:idx]
			} else {
				cur = strs[0][idx]
				if cur != strs[i][idx] {
					return strs[0][:idx]
				}
			}
		}
		idx++
	}
}
