// leetcode 557,反转字符串中的单词2

//给定一个字符串 s ，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序

package main

// 在结尾加一个空格
// 当遇到空格时，就将空格前的单词反转一下
func reverseWords(s string) string {
	first := 0
	end := 0
	flag := false
	tmps := []byte(" " + s + " ")
	for i := 1; i < len(tmps)-1; i++ {
		if tmps[i-1] == ' ' && tmps[i] != ' ' { //开头单词
			first = i
		}
		if tmps[i] != ' ' && tmps[i+1] == ' ' { //结尾单词
			end = i
			flag = !flag
		}
		if flag {
			for j := first; j < (first+end)/2; j++ {
				tmps[j], tmps[first+end-j] = tmps[first+end-j], tmps[j]
			}
			flag = !flag
		}
	}
	return string(tmps[1 : len(tmps)-1])
}

func main() {
	s := "Let's take LeetCode contest"

	reverseWords(s)
}
