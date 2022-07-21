// leetcode 942,增减字符串匹配

package main

// 如果 perm[i] < perm[i + 1] ，那么 s[i] == 'I'
// 如果 perm[i] > perm[i + 1] ，那么 s[i] == 'D'
func diStringMatch(s string) []int {
	res := []int{}
	D, I := len(s), 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			res = append(res, I)
			I++
		} else {
			res = append(res, D)
			D--
		}
	}
	res = append(res, I)
	return res
}
