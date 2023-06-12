// https://leetcode.cn/problems/multiply-strings/
// 43. 字符串相乘
package main

import "strconv"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := make([]int, len(num1)+len(num2))
	j := len(num2) - 1
	index := len(res) - 1
	for j > -1 {
		upset := 0
		curindex := index
		for i := len(num1) - 1; i > -1; i-- {
			cur := int(num1[i]-'0')*int(num2[j]-'0') + upset + res[curindex]
			upset = cur / 10
			res[curindex] = cur % 10
			curindex--
		}
		for upset != 0 {
			res[curindex] = upset % 10
			upset = upset / 10
			curindex--
		}
		index--
		j--
	}
	ans := ""
	for i := 0; i < len(res); i++ {
		if i == 0 && res[i] == 0 {
			continue
		}
		ans += strconv.Itoa(res[i])
	}
	return ans
}
