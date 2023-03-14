package main

// 核心思想：数组其实是一个简易hash表
func isAnagram(s string, t string) bool {
	wordMap := make([]int, 26, 26)
	for _, v := range s {
		wordMap[v-'a']++
	}
	for _, v := range t {
		wordMap[v-'a']--
	}
	for _, v := range wordMap {
		if v != 0 {
			return false
		}
	}
	return true
}
