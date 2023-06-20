// https://leetcode.cn/problems/decode-string/
// 394. 字符串解码
package main

import "strings"

var (
	src string
	ptr int
)

func decodeString(s string) string {
	src = s
	ptr = 0
	return getString()
}
func getString() string {
	if ptr == len(src) || src[ptr] == ']' {
		return ""
	}
	cur := src[ptr]
	repTime := 1
	ret := ""
	if cur >= '0' && cur <= '9' {
		repTime = getDigits()
		ptr++ //此时下标在[
		str := getString()
		ptr++ // 此时下标在]
		ret = strings.Repeat(str, repTime)
	} else if (cur >= 'a' && cur <= 'z') || (cur >= 'A' && cur <= 'Z') {
		ret = string(cur)
		ptr++
	}
	return ret + getString()
}

func getDigits() int {
	ret := 0
	for ; src[ptr] >= '0' && src[ptr] <= '9'; ptr++ {
		ret = ret*10 + int(src[ptr]-'0')
	}
	return ret
}

func main() {
	decodeString("abc2[bc5[de]hg7[oi]l]oing")
}
