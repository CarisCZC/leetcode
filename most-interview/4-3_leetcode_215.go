// https://leetcode.cn/problems/kth-largest-element-in-an-array/
// 215. 数组中的第K个最大元素
package main

// 利用快速排序思想确定第k个最大元素的位置，k~n以及1~k-1 都不需要排序
func findKthLargest(nums []int, k int) int {
	TopKSplit(nums, 0, len(nums)-1, len(nums)-k)
	return nums[len(nums)-k]
}

func TopKSplit(a []int, l, r, index int) {
	if l < r {
		cur := parition(a, l, r)
		if cur == index {
			return
		} else if cur < index {
			TopKSplit(a, cur+1, r, index)
		} else {
			TopKSplit(a, l, cur-1, index)
		}
	}

}

func parition(a []int, l, r int) int {
	if l >= r {
		return -1
	}
	pos := a[l]
	for l < r {
		for l < r && a[r] >= pos {
			r--
		}
		// 此时 a[r] < pos,将a[r]移到左边
		a[l] = a[r]
		for l < r && a[l] < pos {
			l++
		}
		// 同理，a[l] 移到右边
		a[r] = a[l]
	}
	// 到此l==r，两边划分完毕，l的位置就是pos位置
	a[l] = pos
	return l
}
