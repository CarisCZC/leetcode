// leetcode 290,单词规律
// 根据pattern和s是否有相同的规律

// 输入: pattern = "abba", str = "dog cat cat dog"
// 输出: true

package main

import "strings"

// 提供一个pattern中的a->str中的b的映射。
// map
func wordPattern(pattern string, s string) bool {
	tmpmap := map[byte]string{}
	exitmap := map[string]bool{}
	tmps := strings.Split(s, " ")
	if len(pattern) != len(tmps) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		if v, ok := tmpmap[pattern[i]]; !ok { //说明没有建立链接
			if exitmap[tmps[i]] { //pattern[i] 没有出现过，但tmps[i]已经出现过
				return false
			}
			tmpmap[pattern[i]] = tmps[i]
			exitmap[tmps[i]] = true
		} else {
			if tmps[i] != v {
				return false
			}
		}
	}
	return true
}