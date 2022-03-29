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

func generate(numRows int) (nums [][]int) {
	nums = make([][]int, numRows)
	nums[0] = []int{1}
	if numRows >= 2 {
		nums[1] = []int{1, 1}
	}
	for i := 2; i < numRows; i++ {
		nums[i] = generateSan(nums[i-1])
	}
	return

}
