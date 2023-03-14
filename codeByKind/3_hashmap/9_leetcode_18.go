package main

import (
	"sort"
)

// 和15有异曲同工之妙
func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		// if nums[i] > target {
		// 	return res
		// }
		n1 := nums[i]
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			n2 := nums[j]
			if j > i+1 && n2 == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for right > left {
				n3 := nums[left]
				n4 := nums[right]
				tmp := n1 + n2 + n3 + n4
				if tmp < target {
					left++
				} else if tmp > target {
					right--
				} else {
					res = append(res, []int{n1, n2, n3, n4})
					for right > left && nums[left] == nums[left+1] {
						left++
					}
					for right > left && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				}
			}
		}
	}
	return res
}
