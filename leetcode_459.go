//leetcode 459, 重复的子字符串

// 匹配的子字符串逐渐递增
// "abab"
// 先用"a"去比配，不行，用增量"ab"，再匹配

// 最优解：KMP算法，学一学

package main

func repeatedSubstringPattern(s string) bool {
	sublen := 1
	subs := s[:sublen]
next:
	for sublen <= len(s)/2 { //超过一半，就没有匹配的必要了
		for i := 0; i < len(s); i += sublen {
			if s[i:min(i+sublen, len(s))] != subs { //匹配不成功
				sublen++
				subs = s[:sublen]
				continue next
			}
		}
		// 上面循环结束，说明已经匹配完成
		return true
	}
	return false
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
