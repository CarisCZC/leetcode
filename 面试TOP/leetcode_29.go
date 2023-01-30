// https://leetcode.cn/problems/divide-two-integers/?favorite=2ckc81c
package main

import "math"

// 算法复杂度O(n)，但超时，对该题的评价是：无聊
func divide(dividend int, divisor int) int {
	// 判断除数被除数是否同号
	max := math.MaxInt32
	min := math.MinInt32
	sign := true
	p, q := dividend, divisor
	if dividend < 0 {
		sign = !sign
		p = -p
	}
	if divisor < 0 {
		sign = !sign
		q = -q
	}
	res := 0
	if p == 0 {
		return res
	}

	if q != 1 {
		for p >= q {
			i := 1
			tmp := q
			for p >= tmp {
				p -= tmp
				res += i
				i = i << 1
				tmp = tmp << 1
			}

		}
	} else {
		res = p
	}
	if !sign {
		res = -res
	}
	if res > max {
		res = max
	}
	if res < min {
		res = min
	}
	return res
}

func main() {
	divide(32, 5)
}
