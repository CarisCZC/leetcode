// leetcode 844, 比较含退格的字符串

package main

//搞个比较栈
func backspaceCompare(s string, t string) bool {
	res := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			res = append(res, s[i])
		} else {
			if len(res) > 0 {
				res = res[:len(res)-1] //出栈一个
			}
		}
	}
	n := len(res) - 1
	cnt := 0
	for i := len(t) - 1; i > -1; i-- {
		if t[i] != '#' {
			if cnt > 0 {
				// i--
				cnt--
				continue
			}
			if n < 0 {
				res = append(res, t[i])
                n++
			} else if res[n] == t[i] {
				n-- //相当于出栈
			} else {
				return false
			}
		} else {
			cnt++
			continue
		}
	}
	return n < 0
}
