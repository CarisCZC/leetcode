// leetcode 205 同构字符串
// s = "egg", t = "add" 即为同构字符串
// 也就是说，可以确定一个映射关系，在这里，e映射到a，g映射到d
package main

import "strings"

// map创建映射
// 相同位置的只能有相同的映射
func isIsomorphic(s string, t string) bool {
	tmpmap := map[byte]byte{s[0]: t[0]} //单向的映射

	for i := 1; i < len(s); i++ {
		if tmpmap[s[i]] == 0 { //说明未建立映射
			if strings.Contains(t[:i], string(t[i])) {
				// 说明，在i之前，t[i]字符就已经出现过，那么一定有c于t[i]建立过映射
				// 此时，s[i]是无法映射到t[i]的
				return false
			}
			tmpmap[s[i]] = t[i]
			// tmpmap[t[i]] = s[i]
		}
		if tmpmap[s[i]] != t[i] {
			return false
		}
	}
	return true
}
