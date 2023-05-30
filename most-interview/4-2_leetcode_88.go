// https://leetcode.cn/problems/merge-sorted-array/
// 88. 合并两个有序数组

package main

// 从后往前去覆盖
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	idx := len(nums1) - 1
	for i > -1 && j > -1 {
		if nums1[i] > nums2[j] {
			nums1[idx] = nums1[i]
			i--
			idx--
		} else {
			nums1[idx] = nums2[j]
			j--
			idx--
		}
	}
	for i > -1 {
		nums1[idx] = nums1[i]
		i--
		idx--
	}
	for j > -1 {
		nums1[idx] = nums2[j]
		j--
		idx--
	}
}
