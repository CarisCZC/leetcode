// https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations/description/
package main

import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	res := 0
	firstPos := -1
	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			res -= nums[i]
			k--
		} else {
			res += nums[i]
			if firstPos < 0 {
				firstPos = i
			}
		}
	}
	// 此时k还没用完
	if k > 0 && k%2 != 0 { //选择绝对值最小的
		if firstPos == -1 { //说明是一个负数数列,排序后最后一个值最小
			res += nums[len(nums)-1] * 2
		}
		if firstPos == 0 { //说明是一个正数数列，排序后第一个值最小
			res -= nums[0] * 2
		}
		if firstPos > 0 {
			minNeg := -nums[firstPos-1]
			minPos := nums[firstPos]
			if minNeg > minPos {
				res -= minPos * 2
			} else {
				res -= minNeg * 2
			}
		}
	}
	return res
}

func main() {
	largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2)
}
