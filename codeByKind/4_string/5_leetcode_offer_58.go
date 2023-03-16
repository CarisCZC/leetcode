package main

// 左旋字符串
// 就是将前几个字符移到后面
// 全逆，然后0~n-k-1翻转一遍，n-k~n翻转一遍
func reverseLeftWords(s string, n int) string {
	res := []byte(s)
	left, right := 0, len(s)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	// 翻转前n-k
	left, right = 0, len(res)-n-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	// 翻转后k
	left, right = len(res)-n, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return string(res)
}
