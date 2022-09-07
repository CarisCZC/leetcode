// leetcode 806,写字符串需要的行数

package main

// 单词不能被拆开写
func numberOfLines(widths []int, s string) []int {
	res := 0
	for _, v := range s {
		if (res%100 + widths[v-'a']) > 100 {
			res = res + 100 - res%100 + widths[v-'a']
		} else {
			res += widths[v-'a']
		}
	}
	return []int{(res / 100) + 1, res % 100}
}
