// leetcode 349, 两个数组的交集
// 用一个set

package main

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	for _, v := range nums1 {
		if !set[v] {
			set[v] = true
		}

	}
	tmplen := 0
	for _, v := range nums2 {
		if set[v] {
			nums1[tmplen] = v
			set[v] = false
			tmplen++
		}
	}
	return nums1[:tmplen]

}
