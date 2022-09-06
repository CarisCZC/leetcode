// leetcode 219, 存在重复元素Ⅱ

package main

func containsNearbyDuplicate(nums []int, k int) bool {

	set := make(map[int]int) //内部int表示坐标差

	for i, v := range nums {
		if index, ok := set[v]; ok { //已经存在过了，此时index是该值第一次出现时的坐标
			if i-index <= k {
				return true
			}
			// 	} else {
			// 		set[v] = i //舍弃第一次出现的值，存储该次坐标
			// 	}
			// } else {
			// 	set[v] = i
			// }
		}
		// 优化
		set[v] = i
	}

	return false
}
