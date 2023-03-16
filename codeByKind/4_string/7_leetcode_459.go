package main

// 459.重复的字符串
// s+s 中间一定有一个重复的s
func repeatedSubstringPattern(s string) bool {
	newS := s + s
	n := len(s)
	for i := 1; i < len(newS)-n; i++ {
		tmp := newS[i : i+n]
		if tmp == s {
			return true
		}
	}
	return false
}
