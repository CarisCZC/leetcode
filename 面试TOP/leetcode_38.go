// https://leetcode.cn/problems/count-and-say/?favorite=2ckc81c
package main

import (
	"strconv"
)

func countAndSay(n int) string {
	res := "1"
	if n == 1 {
		return "1"
	}
	for i := 1; i < n; i++ {
		res = discribe(res)
	}
	return res
}

func discribe(input string) string {
	output := ""
	flag := input[0]
	cnt := 0
	for i := 0; i < len(input); i++ {
		if input[i] == flag {
			cnt++
		} else {
			output += strconv.Itoa(cnt) + string(flag)
			flag = input[i]
			cnt = 1
		}
		// 最后一个字符
		if i == len(input)-1 {
			output += strconv.Itoa(cnt) + string(flag)
		}
	}
	return output
}
func main() {
	countAndSay(4)
}
