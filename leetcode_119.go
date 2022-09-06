//leetcode 119, 杨辉三角2
package main

func generateSan(lastLevel []int) (thisLevel []int) {
	thisLen := len(lastLevel) + 1
	thisLevel = make([]int, thisLen)
	thisLevel[0] = 1
	thisLevel[thisLen-1] = 1
	for i := 1; i < thisLen-1; i++ {
		thisLevel[i] = lastLevel[i] + lastLevel[i-1]
	}
	return
}

func getRow(rowIndex int) []int {
	nums := make([]int, rowIndex+1)
	nums[0] = 1
	if rowIndex == 0 {
		return nums
	}
	if rowIndex >= 1 {
		nums[1] = 1
	}
	for i := 2; i <= rowIndex; i++ {

		nums = generateSan(nums[:i])
	}
	return nums
}
