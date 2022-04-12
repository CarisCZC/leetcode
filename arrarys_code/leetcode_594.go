// leetcode 594,最长和谐子序列
package main

// 先排序，然后两头往中间找
func findLHS(nums []int) int {
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v] += 1
	}
	ans := 0
	for key, v := range mp {
		if mp[key+1] == 0 {
			continue
		}
		tmpans := mp[key+1] + v
		if tmpans > ans {
			ans = tmpans
		}
	}
	return ans
}
