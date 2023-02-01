// https://leetcode.cn/problems/group-anagrams/?favorite=2ckc81c
package main

// 要用到map
// 面对字符的题，好好考虑一下排序
func groupAnagrams(strs []string) [][]string {
	group := map[[26]int][]string{}

	for _, str := range strs {
		char := [26]int{}
		// 统计字符情况
		for _, ch := range str {
			char[ch-'a']++
		}
		group[char] = append(group[char], str)
	}
	ans := make([][]string, 0, len(group))
	for _, v := range group {
		ans = append(ans, v)
	}
	return ans
}
