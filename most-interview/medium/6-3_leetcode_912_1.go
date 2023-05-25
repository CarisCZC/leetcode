package main

// 利用map实现桶排序
func sortArray(nums []int) []int {
	myMap := make(map[int]int, 100001) // 因为有0
	for _, v := range nums {
		myMap[v+50000]++
	}
	i := 0
	key := 0
	for key < 100001 {
		if count, ok := myMap[key]; ok {
			for j := 0; j < count; j++ {
				nums[i] = key - 50000
				i++
			}
		}
		key++
	}
	return nums
}
