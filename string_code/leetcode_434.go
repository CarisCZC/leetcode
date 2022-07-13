// leetcode 434,字符串中的单词数
// 这里的单词指的是连续的不是空格的字符

package main

// 找到第一个非空格，然后确定下一个空格的位置。这两个之间是一个单词
// 然后从下一个空格开始，找下一个非空格的位置，以及下下个非空格
func countSegments(s string) int {
	tmps := s + " " //尾部添加空格，方便处理尾部单词
	res := 0
	flag := false // true 表示非空格已经找到，找下一个空格
	for i := 0; i < len(tmps); i++ {
		if !flag && tmps[i] != ' ' { //非空格
			flag = !flag
		}
		if flag && tmps[i] == ' ' { // 非空格已经找到，且当前为是空格
			flag = !flag
			res++
		}
	}
	return res
}
