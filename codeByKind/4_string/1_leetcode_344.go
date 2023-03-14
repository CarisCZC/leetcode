package main

// 反转字符串
// 原地修改数组
func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		tmp := s[i]
		s[i] = s[j]
		s[j] = tmp
		i++
		j--
	}
}
