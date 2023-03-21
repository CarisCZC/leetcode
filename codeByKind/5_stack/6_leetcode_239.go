// https://leetcode.cn/problems/sliding-window-maximum/
package main

// 维护一个单调队列
func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	// 维护的是一个单调递减队列，q中存放的是下标，下标按递增顺序
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		// 如果下标小于当前滑动窗口的范围，即删除这个下标。
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}
