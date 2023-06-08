// https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof/
// 剑指 Offer 40. 最小的k个数
package main

import "math/rand"

// 和另一个题，最大的k个数思路是一致的
func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	return partition(arr, k)
}

func partition(arr []int, k int) []int {
	index := rand.Intn(len(arr))
	part := arr[index]
	left, right := 0, len(arr)-1
	arr[left], arr[index] = arr[index], arr[left]
	for left < right {
		// 一定要先让right往left这边移动，减少出BUG的可能性
		for left < right && arr[right] >= part {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] < part {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = part
	res := arr[:left+1]
	if left == k-1 {
		return res
	}
	if left > k-1 {
		res = partition(res, k)
	} else {
		res = append(res, partition(arr[left+1:], k-left-1)...)
	}
	return res
}
