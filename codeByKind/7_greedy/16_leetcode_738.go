package main

import "strconv"

func monotoneIncreasingDigits(n int) int {
	// 如果n不是单调递增的，那么，n = 递增部分的最小的那个数+999...（不递增的部分都为9）
	if n < 10 {
		return n
	}
	// 转为string,好操作
	nStr := strconv.Itoa(n)
	res := []byte(nStr)
	for i := len(res) - 1; i > 0; i-- {
		if res[i] < res[i-1] { //非递增
			for j := i; j < len(res); j++ {
				res[j] = '9'
			}
			res[i-1]-- // 后面的变成99,前一位一定要-1，这样才能保证此时res<n
		}
	}
	num, _ := strconv.Atoi(string(res))
	return num
}
