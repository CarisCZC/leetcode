// leetcode 596, 下一个更大元素 Ⅰ
package main

// 开辟set，set[v]表示nums2中v的下一个更大元素

// 极限情况下，即nums2是递减序列时，时间复杂度为O(n^2)
// 主要开销在如何找下一个最大值上
// func nextGreaterElement(nums1 []int, nums2 []int) (ans []int) {
// 	nextMaxMap := make(map[int]int, len(nums2))
// 	ans = make([]int, 0, len(nums1))
// 	for i := 0; i < len(nums2); i++ {
// 		nextMaxMap[nums2[i]] = -1
// 		for j := i + 1; j < len(nums2); j++ {
// 			if nums2[j] > nums2[i] {
// 				nextMaxMap[nums2[i]] = nums2[j]
// 				break
// 			}
// 		}
// 	}
// 	for _, v := range nums1 {
// 		ans = append(ans, nextMaxMap[v])
// 	}
// 	return
// }

// 单调栈写法
// 减少了找下一个最大值的时间
func nextGreaterElement(nums1 []int, nums2 []int) (res []int) {
	ansMap := make(map[int]int, len(nums2))
	stack := []int{}
	for i := len(nums2) - 1; i > -1; i-- {
		num := nums2[i]
		for len(stack) > 0 && num >= stack[len(stack)-1] { //num大于栈顶元素
			stack = stack[:len(stack)-1] //出栈
		}
		if len(stack) > 0 {
			ansMap[num] = stack[len(stack)-1]
		} else {
			ansMap[num] = -1
		}
		stack = append(stack, num)
	}
	res = make([]int, 0, len(nums1))
	for i, v := range nums1 {
		res[i] = ansMap[v]
	}
	return
}
