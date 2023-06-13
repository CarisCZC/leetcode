// https://leetcode.cn/problems/reverse-words-in-a-string/
// 151. 反转字符串中的单词
package main

// 双指针去空格
func reverseWords(s string) string {
	head, tail := 0, len(s)-1
	sbyte := []byte(s)
	for head < tail && sbyte[head] == ' ' {
		head++
	}
	for head < tail && sbyte[tail] == ' ' {
		tail--
	}
	// 整个字符串逆转
	sbyte = sbyte[head : tail+1]
	sbyte = reverse(sbyte, 0, len(sbyte)-1)
	// 对每个单词逆转
	for i := 0; i < len(sbyte); i++ {
		if sbyte[i] == ' ' {
			continue
		}
		idx := i
		for idx < len(sbyte) && sbyte[idx] != ' ' {
			idx++
		}
		sbyte = reverse(sbyte, i, idx-1)
		i = idx
	}
	// 去除单词中间的空格
	head = 0
	tail = 0
	for tail < len(sbyte) {
		if tail != 0 && sbyte[tail] == ' ' && sbyte[tail-1] == ' ' {
			tail++
			continue
		}
		sbyte[head] = sbyte[tail]
		head++
		tail++
	}
	ans := string(sbyte[:head])
	return ans
}

func reverse(s []byte, start, end int) []byte {
	for i, j := start, end; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}
