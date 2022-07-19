// leetcode 830,较大分组的位置

package main

//在一个由小写字母构成的字符串 s 中，包含由一些连续的相同字符所构成的分组。
func largeGroupPositions(s string) [][]int {
	res := [][]int{}
	tmp := s[0]
	cnt := 1
	for i := 1; i < len(s); i++ {
		if s[i] == tmp {
			cnt++
		} else {
			if cnt >= 3 {
				res = append(res, []int{i - cnt, i - 1})
			}
			tmp = s[i]
			cnt = 1
		}
	}
	if cnt >= 3 {
		res = append(res, []int{len(s) - cnt, len(s) - 1})
	}
	return res
}
