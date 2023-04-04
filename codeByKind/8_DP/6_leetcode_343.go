package main

// 尽可能多的分
// 全分成1不行，全分成2
func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	k := 2 // 至少分成两个
	res := 0
	for n/k >= 1 {
		part := n / k
		other := 0
		tmp := 1
		if n%k != 0 {
			other = n % k
		}
		for i := 0; i < k; i++ {
			if i < k-other {
				tmp *= part
			} else {
				tmp *= (part + 1)
			}
		}
		res = max(res, tmp)
		k++
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
