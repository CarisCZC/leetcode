package main

// 要求递增
// 要用used数组来去重
func findSubsequences(nums []int) [][]int {
	path := []int{}
	res := [][]int{}
	var backtarck func(nums []int, start int)
	backtarck = func(nums []int, start int) {
		if len(path) > 1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		used := make(map[int]bool, len(nums)) // 对同一层元素去重
		for i := start; i < len(nums); i++ {
			// 去重
			if used[nums[i]] {
				continue
			}
			// if i != 0 && nums[i-1] == nums[i] { // 因为这道题没法排序，这种去重方式不可用
			// 	continue
			// }
			if len(path) == 0 || nums[i] >= path[len(path)-1] {

				path = append(path, nums[i])
				used[nums[i]] = true
				backtarck(nums, i+1)
				path = path[:len(path)-1]
			}
		}
	}
	backtarck(nums, 0)
	return res
}
