package main

func longestConsecutive(nums []int) int {
	numsMap := map[int]bool{}
	for _, v := range nums {
		numsMap[v] = true
	}
	res := 0
	for num := range numsMap {
		if !numsMap[num-1] {
			currentNum := num
			currentLen := 1
			for numsMap[currentNum+1] {
				currentNum++
				currentLen++
			}
			if currentLen > res {
				res = currentLen
			}
		}
	}
	return res
}
