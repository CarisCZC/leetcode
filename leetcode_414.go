// leetcode 414, 第三大的数
package main

// 空间换时间法
// func thirdMax(nums []int) int {
// 	numsSet := make(map[int]bool)
// 	for _, v := range nums {
// 		if !numsSet[v] {
// 			numsSet[v] = true
// 		}
// 	}
// 	keys := make([]int, 0, len(numsSet))
// 	for k := range numsSet {
// 		keys = append(keys, k)
// 	}
// 	sort.Ints(keys)
// 	if len(keys) >= 3 {
// 		return keys[len(keys)-3]
// 	}
// 	return keys[len(keys)-1]

// }

// 一次遍历法
type maxNums struct {
	num  int
	flag bool
}

func thirdMax(nums []int) int {
	first := maxNums{0, false}
	second := maxNums{0, false}
	ans := maxNums{0, false}
	for _, v := range nums {
		if !first.flag || v > first.num {
			ans, second, first = second, first, maxNums{v, true}
		} else if v < first.num && (!second.flag || v > second.num) {
			ans, second = second, maxNums{v, true}
		} else if v < second.num && (!ans.flag || v > ans.num) {
			ans = maxNums{v, true}
		}
	}
	if !ans.flag {
		return first.num
	}
	return ans.num
}
