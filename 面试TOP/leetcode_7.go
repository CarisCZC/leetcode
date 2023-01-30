// https: //leetcode.cn/problems/reverse-integer/?favorite=2ckc81c
package main

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	// 转为字符串
	xString := strconv.Itoa(x)
	newNum := ""
	for i := len(xString) - 1; i > 0; i-- {
		newNum += xString[i : i+1]
	}
	if xString[0] != '-' {
		newNum += xString[:1]
	} else {
		newNum = xString[:1] + newNum
	}
	res, _ := strconv.Atoi(newNum)
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}
