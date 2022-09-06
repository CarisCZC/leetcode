// leetcode 242, 有效的字母异位词
// 若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
package main

// 一个map
func isAnagram(s string, t string) bool {
	tmpmap := map[byte]int{}
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		tmpmap[s[i]]++
		tmpmap[t[i]]--
	}
	for _, v := range tmpmap {
		if v != 0 {
			return false
		}
	}
	return true
}
