// https://leetcode.cn/problems/valid-parentheses/
package main

func isValid(s string) bool {
	kuohao := []byte{}
	offset := 0
	smap := map[byte]byte{'[': ']', '{': '}', '(': ')'}
	for i := 0; i < len(s); i++ {
		kuohao = append(kuohao, s[i])
		offset++
		if offset > 1 && kuohao[offset-1] == smap[kuohao[offset-2]] {
			kuohao = kuohao[:offset-2]
			offset -= 2
		}
	}
	if offset != 0 {
		return false
	}
	return true
}

func main() {
	isValid("()[]{}")
}
