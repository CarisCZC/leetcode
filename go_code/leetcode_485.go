// leetcode 485, 最大连续1的个数
package main

func findMaxConsecutiveOnes(nums []int) int {
	ans := 0
	tmpCot := 0
	for _, v := range nums {
		if v == 1 {
			tmpCot++
		} else {
			if tmpCot > ans {
				ans = tmpCot

			}
			tmpCot = 0
		}
	}
	if tmpCot > ans {
		return tmpCot
	}
	return ans
}
