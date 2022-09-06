// leetcode 645,错误的集合
package main

//普遍做法：创建一个set，来存储已经查过的，当有重复时，返回即可
//空间O(1)做法？
// 先排序，然后查找那个错误值

// 创建一个固定长度的数组，然后按照数值插入相应的坐标
func findErrorNums(nums []int) (ans []int) {
	ans = make([]int, 0, 2)
	tmp := make([]bool, len(nums))
	for _, v := range nums {
		if !tmp[v-1] {
			tmp[v-1] = true
		} else {
			ans = append(ans, v)
		}
	}
	for i, v := range tmp {
		if !v {
			ans = append(ans, i+1)
			break
		}
	}
	return ans
}
