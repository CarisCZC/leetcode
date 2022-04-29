// leetcode 33,搜索旋转排序数组

package main

// 首先来说，nums按升序排列，且在某个位置上进行了旋转，而且每个数都不会相同

// [4,5,6,7,0,1,2,3] 找0，3大于0，说明从后往前最多找3-0=3位即可
// [4,5,6,7,0,1,2,3] 找5，3小于5，则用5-4=1，在4往后最多找1位

// 算法复杂度，最差情况是查找的是极值时，取决于边界到极值的最小值
// 虽然能AC，但是算法并没有完美实现
// 因为当是升序时，默认就会从右往左找，此时是O(n)!
func search(nums []int, target int) int {
	r := len(nums) - 1
	if target > nums[r] { //此时从左边开始找
		fid := target - nums[0]
		for i := 0; i <= r && i <= fid; i++ {
			if nums[i] > target {
				break
			} else if nums[i] == target {
				return i
			}
		}
	} else {
		fid := nums[r] - target
		for i := 0; r-i > -1 && i <= fid; i++ {
			if nums[r-i] < target {
				break
			} else if nums[r-i] == target {
				return r - i
			}
		}
	}
	return -1
}
