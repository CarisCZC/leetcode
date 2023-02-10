// https://leetcode.cn/problems/minimum-window-substring/?favorite=2ckc81c
package main

import "math"

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	need, window := map[byte]int{}, map[byte]int{}
	left, right, valid := 0, 0, 0
	start, length := 0, math.MaxInt32
	for i := range t {
		need[t[i]]++
	}
	for right < len(s) {
		c := s[right]
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// 即当前窗口里已经包含了一个t，考虑窗口收缩
		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}
