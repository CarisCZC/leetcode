// leetcode 168. Excel表列名称
// A -> 1,Z -> 26,AA -> 27
package main

import "fmt"

// 其实就是26进制
// ASCII // A-Z 65-90
func convertToTitle(columnNumber int) string {
	//除26，找对应的编码
	res := ""
	n := columnNumber //除数
	m := n % 26       //余数
	for n > 0 {
		if n < 27 { //说明此时n只是一位，直接表示即可
			res = fmt.Sprint(n+64) + res
			break
		}
		if m == 0 {
			m = 26
			n--
		}
		res = fmt.Sprint(m+64) + res
		n = n / 26
		m = n % 26

	}
	return res
}

func main() {
	n := 701
	convertToTitle(n)
}
