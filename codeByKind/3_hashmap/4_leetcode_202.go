package main

// 主要判断这个数是否会再出现，即陷入循环
func isHappy(n int) bool {
	resSet := map[int]bool{}
	for n != 1 && !resSet[n] {
		resSet[n] = true
		tmp := n
		sum := (tmp % 10) * (tmp % 10)
		for tmp > 0 {
			tmp /= 10
			sum += (tmp % 10) * (tmp % 10)
		}
		n = sum
	}
	return n == 1
}
