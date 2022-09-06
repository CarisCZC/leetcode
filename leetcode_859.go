// leetcode 859, 亲密字符串

package main

// s和goal只能由两个位置的字符不一样，且这两个字符交叉一样。
// 或者，没有字符不一样
func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) || len(s) == 1 {
		return false
	}
	pre, this := -1, -1
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			if pre == -1 {
				pre = i
			} else if this == -1 {
				this = i
			} else { // 有第三个不一样
				return false
			}
		}
	}
	// 如果两个字符串一样：
	// 查字符串有没有一样的字符出现过两次？
	if pre == -1 && this == -1 {
		tmp := map[rune]int{}
		for _, v := range s {
			tmp[v]++
			if tmp[v] >= 2 {
				return true
			}
		}
		return false
	}
	if pre != -1 && this != -1 && s[pre] == goal[this] && s[this] == goal[pre] {
		return true
	}
	return false
}
