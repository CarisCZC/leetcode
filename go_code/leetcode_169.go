//leetcode 169,多数元素
// 常规解法，开hash
// 摩根投票法
// 目标数target>[n/2],也就是说，比其他任意数出现的总和都要多，因此就1对1来消除。最后我总会剩下一个
package main

func majorityElement(nums []int) int {
	ans := nums[0]
	count := 0
	for _, v := range nums {
		if v != ans {
			if count == 0 {
				ans = v
				count++
			} else {
				count--
			}
		} else {
			count++
		}

	}
	return ans

}
