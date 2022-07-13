//leetcode 344， 反转字符串

package main

// 输入是一个字符数组，将其反转

// 双指针，前后交换
func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	for i < j {
		tmp := s[i]
		s[i] = s[j]
		s[j] = tmp
		i++
		j--
	}
}
