// leetcode 674,最长连续递增子数列
package main

//
func findLengthOfLCIS(nums []int) int {
	ans := 1
	tmpAns := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			tmpAns++
		} else { //此时nums[i+1] <= nums[i]，也就是前面的连续断开了
			// 这时候只需要比较即可。
			if tmpAns > ans {
				ans = tmpAns
			}
			tmpAns = 1 //重置tmpAns，此时i从下个位置重新算起
		}
	}
	if tmpAns > ans {
		return tmpAns
	}
	return ans
}
