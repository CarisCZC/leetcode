// leetcode 15, 三数之和, mid
package main

import "sort"

// 找出所有和为0，且不重复的三元组
// 三元组的情况：两正一负，一正一负和0，两负一正
// 排序+双指针
// 去除重复，比较麻烦
// 这里的想法是：对于一个确定的数a,找到其所有可能的解法，如果在遇到a，直接跳过
func threeSum(nums []int) [][]int {
	res := make([][]int, 0, len(nums))
	sort.Ints(nums)

	if len(nums) < 3 {
		return res
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] { //去除重复解
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			tmp := nums[i] + n2 + n3
			if tmp == 0 {
				res = append(res, []int{nums[i], n2, n3})
				for l < r && nums[r] == n3 {
					r--
				}
				for l < r && nums[l] == n2 {
					l++
				}
			} else if tmp < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, -1}
	threeSum(nums)
}
