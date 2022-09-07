// leetcode 58,最后一个单词的长度
package main

// 返回字符串中的最后一个单词的长度，单词之间有不止一个空格，

//"luffy is still joyboy"

// 从后往前遍历，先找到最后一个非空格，记录其下标a
// 然后继续往前找，找到距离a最近的空格，下标b
// 单词长度，a-b
func lengthOfLastWord(s string) int {
	a := len(s)
	b := -1
	for i := a - 1; i > -1; i-- {
		if a == len(s) && s[i] != ' ' { // a == len(s)确保a只赋值一次
			a = i
		}
		if a != len(s) && s[i] == ' ' { // a != len(s)确保a已经赋值了
			b = i
			break
		}

	}
	return a - b
}
