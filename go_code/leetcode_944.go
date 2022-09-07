// leetcode 944 ,删列造序

package main

// 删除某个不是递增的列
func minDeletionSize(strs []string) int {
	res := 0
next:
	for i := 0; i < len(strs[0]); i++ {
		flag := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if min(flag, strs[j][i]) {
				// 记录当前前面的最大
				flag = strs[j][i]
			} else {
				res++
				continue next
			}
		}
	}
	return res
}

func min(a, b byte) bool {
	if a <= b {
		return true
	}
	return false
}
