// https://leetcode.cn/problems/h-index/?envType=study-plan-v2&id=top-interview-150
// 274. H 指数
package main

// 就是快排思想
// 直接排序
// 最优解法是计数排序
func hIndex(citations []int) int {
	n := len(citations)
	cnt := make([]int, n+1) //cnt[i] 表示引用>=i的论文有几篇
	for _, cncitation := range citations {
		if cncitation >= n {
			cnt[n]++
		} else {
			cnt[cncitation]++
		}
	}
	for i, tot := n, 0; i >= 0; i-- {
		tot += cnt[i]
		if tot >= i {
			return i
		}
	}
	return 0
}
