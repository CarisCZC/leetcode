package main

// 和1，2思路均类似
func intersection(nums1 []int, nums2 []int) []int {
	numsMap := map[int]bool{}
	res := []int{}
	for _, num := range nums1 {
		numsMap[num] = true
	}
	for _, num := range nums2 {
		if numsMap[num] {
			res = append(res, num)
			numsMap[num] = false
		}
	}
	return res
}
