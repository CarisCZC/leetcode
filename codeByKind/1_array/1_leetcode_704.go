// 二分查找
package main

// 二分查找分为两种
// 1. 区间为左闭右闭[i,j]，此时left<=right，right=middle-1

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		middle := (l + r) / 2
		cur := nums[middle]
		if cur == target {
			return middle
		}
		if cur > target { // 往左区间找
			r = middle - 1
		} else { //往右区间找
			l = middle + 1
		}
	}
	return -1
}

// 2. 区间为左闭右开，此时left<right，right=middle
func search2(nums []int, target int) int {
	// 因为右边界不被包括，所以直接是nums的长度
	l, r := 0, len(nums)
	mid, cur := -1, -1
	for l < r {
		mid = l + (r-l)/2
		cur = nums[mid]
		if cur == target {
			return mid
		} else if cur > target {
			//往右区间找
			r = mid
		} else {
			l = mid + 1
		}

	}
	return mid
}
