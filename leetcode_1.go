package main

import (
	"sort"
)

// leetcode 题号：1 两数之和

// 很烂的算法
func twoSum(nums []int, target int) []int {
	// 首先给数组排序
	tmpNums := make([]int,len(nums))
	copy(tmpNums,nums)
	sort.Ints(tmpNums)
	left :=  0
	right := len(tmpNums)-1
	result := tmpNums[left]+tmpNums[right]

	for ; result != target; result = tmpNums[left]+tmpNums[right]{
		// 如果和大于目标值，右值-1
		if result > target{
			if right > left{
				right--
			}
		}
		// 如果和小于目标值，左值坐标+1
		if result < target{
			if left < right{
				left++
			}
		}
	}
	leftFlag := false
	rightFlag := false
	//得到left和right后去源数组找坐标
	for i,j:=0,len(tmpNums)-1; i<len(tmpNums) && j>=0;{

		if leftFlag && rightFlag{
			break
		}

		if !leftFlag && nums[i] == tmpNums[left] {
			left = i
			leftFlag = true
		}
		if !rightFlag && nums[j] == tmpNums[right] {
			right = j
			rightFlag = true
		}
		if !leftFlag{
			i++
		}
		if !rightFlag{
			j--
		}
	}
	resultNums := []int{left,right}
	return resultNums
}

// 标准答案版本
//作者：himymBen
//链接：https://leetcode-cn.com/problems/two-sum/solution/pythonjavajavasc-by-himymben-o8fx/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func twoSum2(nums []int, target int) (ans []int) {
	idxMap := map[int]int{}
	for i, num := range nums {
		if v, ok := idxMap[target - num]; ok {
			ans = []int{v, i}
			break
		}
		idxMap[num] = i
	}
	return
}

