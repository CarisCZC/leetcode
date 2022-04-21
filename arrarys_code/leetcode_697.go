// leetcode 695, 数组的度

package main

//使用map，记录坐标
func findShortestSubArray(nums []int) int {
	// 这里使用[]int空间占用比较多
	mp := make(map[int][]int, len(nums))
	for i, v := range nums {
		if n, ok := mp[v]; ok {
			n[0]++
			n[2] = i
		} else {
			n := make([]int, 3) //第一个数，存该元素出现的次数，第二个数，存其最短连续子数组长度
			n[0] = 1
			n[1] = i
			n[2] = i
			mp[v] = n
		}
	}
	tmplen := 0
	ans := 0
	for _, v := range mp {
		tmpAns := v[2] - v[1] + 1
		if v[0] > tmplen {
			tmplen = v[0]
			ans = tmpAns
		} else if v[0] == tmplen {
			if tmpAns < ans {
				ans = tmpAns
			}
		}
	}
	return ans
}
