// leetcode 448, 数组中消失的数
package main

// 以自身数组充当哈希表的解法
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	ans := make([]int, 0, len(nums))
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return ans
}
