package main

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	cur := []int{}
	tmpSum := 0
	var backtrack func(k, n, index int)
	backtrack = func(k, n, index int) {
		for len(cur) == k && tmpSum == n {
			tmp := make([]int, k)
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		for i := index; i <= 9; i++ {
			if tmpSum > n {
				break
			}
			cur = append(cur, i)
			tmpSum += i
			backtrack(k, n, i+1)
			tmpSum -= i
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(k, n, 1)
	return res
}
