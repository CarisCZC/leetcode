// https://leetcode.cn/problems/merge-sorted-array/?envType=study-plan-v2&id=top-interview-150
// 1. 合并两个有序数组

package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	tmpNums := append([]int{}, nums1[:m]...)
	index := 0
	i, j := 0, 0
	for i < m && j < n {
		a, b := tmpNums[i], nums2[j]
		if a <= b {
			nums1[index] = a
			i++
		} else {
			nums1[index] = b
			j++
		}
		index++
	}
	for i < m { // 数组1没有比完
		nums1[index] = tmpNums[i]
		i++
		index++
	}
	for j < n { // 数组2没有比完
		nums1[index] = nums2[j]
		j++
		index++
	}
}

// func main() {
// 	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
// }
