// leetcode 520, 检测大写字母

package main

// 要么全大写
// 要么全小写
// 要么只有首字母大写
// 循环走一遍，统计大写字母的个数
func detectCapitalUse(word string) bool {
	cnt := 0
	for i := 0; i < len(word); i++ {
		if word[i] >= 65 && word[i] <= 90 {
			cnt++
		}
	}
	if cnt == 0 {
		return true
	}
	if cnt == len(word) {
		return true
	}
	if cnt == 1 && (word[0] >= 65 && word[0] <= 90) {
		return true
	}

	return false
}
