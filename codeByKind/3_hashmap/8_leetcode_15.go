package main

import "sort"

// hash能做，但复杂。使用双指针，会更快一些
func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for right > left {
			tmp := nums[i] + nums[left] + nums[right]
			if tmp > 0 {
				right--
			} else if tmp < 0 {
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
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
	return res
}
