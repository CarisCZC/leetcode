// leetcode 350, 两个数组的交集Ⅱ
// 349的算法在此依然可以用
package main

func intersect(nums1 []int, nums2 []int) []int {
	set := make(map[int]int)
	num1Len := len(nums1)
	num2Len := len(nums2)
	if num1Len < num2Len {
		for _, v := range nums1 {
			set[v]++
		}  
		tmplen := 0
		for _, v := range nums2 {
			if set[v] > 0 {
				nums1[tmplen] = v
				set[v]--
				tmplen++
			}
		}
		return nums1[:tmplen]
	} else {
		for _, v := range nums2 {
			set[v]++
		}
		tmplen := 0
		for _, v := range nums1 {
			if set[v] > 0 {
				nums2[tmplen] = v
				set[v]--
				tmplen++
			}
		}
		return nums2[:tmplen]
	}
}
