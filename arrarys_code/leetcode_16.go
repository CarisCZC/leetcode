// leetcode 16,最接近的三数之和，mid

package mian

import (
	"math"
	"sort"
)

// 这题和第15题是非常类似的
// 保存当前组合与target的差值，差值最小的存储在res中
func threeSumClosest(nums []int, target int) int {
	res := 0
	diff := math.MaxInt
	// 先排序
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		n1 := nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			tmpres := n1 + n2 + n3
			tmpDiff := 0
			if tmpres == target {
				return target
			}
			if tmpres > target { //说明和比较大，这时，r需要减小
				tmpDiff = tmpres - target
				r--
			} else {
				tmpDiff = target - tmpres
				l++
			}
			if tmpDiff < diff {
				res = tmpres
				diff = tmpDiff
			}
		}
	}
	return res
}
