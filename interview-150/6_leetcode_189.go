// https://leetcode.cn/problems/rotate-array/?envType=study-plan-v2&id=top-interview-150
// 189. 轮转数组

package main

// 先逆置所有元素，然后将后n-k%n个元素翻转一下，
func rotate(nums []int, k int) {

	n := len(nums)
	if n == 1 {
		return
	}
	// 交换逆置
	for i, j := 0, n-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	newK := k % n
	// 后n-k个元素逆置
	for i, j := newK, n-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	// 前k逆置
	for i, j := 0, newK-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 15)
}
