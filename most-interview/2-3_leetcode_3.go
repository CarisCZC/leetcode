// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
// 3. 无重复字符的最长子串

package main

// 滑动窗口题目
// 右指针一直往前找，直到碰到一个重复字符。
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)
	rk, ans := -1, 0         // 右指针
	for i := 0; i < n; i++ { // 左指针就是i
		for rk+1 < n && m[s[rk+1]] == 0 { //如果该字符不重复，就向右滑动，先判断能不能重复与否，在进行滑动
			m[s[rk+1]]++
			rk++
		}
		if rk-i+1 > ans {
			ans = rk - i + 1
		}
		// 左指针开始滑动
		delete(m, s[i])
	}
	return ans
}
