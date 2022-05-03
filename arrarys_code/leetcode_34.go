// leetcode 34，在排序数组中查找元素的第一个和最后一个位置
// 本质上还是二分查找

package main

// 要么能找到，要么没有
// 如果找不到，直接返回[-1,-1]
// 如果能找到：那么找到的必然是其中的一个（不一定是第一个或者最后一个）
// 这时候以此位置开始往两边找（相等的一定在一起）

func searchRange(nums []int, target int) []int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r { // 这里必须是l <= r ，不然在只有一个元素的情况下会出错
		mid := (l + r) >> 1
		tmp := nums[mid]
		if tmp == target { //往左右找
			l = mid
			for l-1 >= 0 && nums[l-1] == target {
				l--
			}
			r = mid
			for r+1 < n && nums[r+1] == target {
				r++
			}
			return []int{l, r}
		} else if tmp > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return []int{-1, -1}
}
