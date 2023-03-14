package main

// 也是map的思想
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	mapA := map[int]int{}
	for _, a := range nums1 {
		for _, b := range nums2 {
			mapA[a+b]++
		}
	}
	cnt := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			if tmp, ok := mapA[0-(c+d)]; ok {
				cnt += tmp
			}
		}
	}
	return cnt
}
