// leetcode 796,旋转字符串

package main

// 环式比较：分别以s[i]为开头组建新字符串和goal比较
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	n := len(s)
next:
	for i := 0; i < n; i++ {
		fi := i
		for k := 0; k < n; k++ {
			if s[fi%n] != goal[k] {
				continue next
			}
			fi++
		}
		return true
	}
	return false
}
