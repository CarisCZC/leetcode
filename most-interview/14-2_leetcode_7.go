// https://leetcode.cn/problems/reverse-integer/
// 7. 整数反转

package main

func reverse(x int) int {
	// 判断正负
	// min := math.MinInt32
	// max := math.MaxInt32
	res := 0
	for x != 0 {
		tmp := x % 10
		if res < -214748364 || (res == -214748364 && tmp < -8) {
			return 0
		}
		if res > 214748364 || (res == 214748364 && tmp > 7) {
			return 0
		}
		res = res*10 + tmp
		x = x / 10
	}
	return res
}

func main() {
	reverse(1534236469)
}
