// https://leetcode.cn/problems/permutations/
// 46. 全排列

package main

func permute(nums []int) [][]int {
	res := [][]int{}
	isUsed := make([]bool, len(nums))

	res = helper(nums, []int{}, res, isUsed, 0)
	return res
}

func helper(nums []int, cur []int, res [][]int, isUsed []bool, deep int) [][]int {
	if len(nums) == deep {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		res = append(res, tmp)
		return res
	}
	for i := range nums {
		if !isUsed[i] {
			isUsed[i] = true
			cur = append(cur, nums[i])
			deep++
			res = helper(nums, cur, res, isUsed, deep)
			cur = cur[:len(cur)-1]
			deep--
			isUsed[i] = false
		}
	}
	return res
}

func main() {
	permute([]int{5, 4, 6, 2})
}
