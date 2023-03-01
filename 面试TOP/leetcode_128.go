package main

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longestLen := 0
	for num := range numSet {
		// 当不连续的时候
		if !numSet[num-1] {
			currentNum := num
			currentLen := 1
			for numSet[currentNum+1] {
				currentNum++
				currentLen++
			}
			if longestLen < currentLen {
				longestLen = currentLen
			}
		}
	}
	return longestLen
}
