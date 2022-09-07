// leetcode 541,反转字符串

// 给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。

// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

package main

// 第一个k反转，第二个k不反转
// 即，反转奇数k
func reverseStr(s string, k int) string {
	str := []byte(s)
	for i := 0; i < len(str); i += 2 * k {
		//i~i+k是需要反转的字符
		n := min(i+k, len(s))
		for j := i; j < (i+n)/2; j++ {
			str[j], str[i+n-j-1] = str[i+n-j-1], str[j]
		}
	}
	return string(str)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := "abcdefg"
	k := 2
	reverseStr(s,k)
}
