// leetcode 27, 搜索插入位置
// 其实就是二分查找
package main

func searchInsert(nums []int, target int) int {
	tmpMin := 0
	tmpMax := len(nums)
	tmpIndex := 0
	for tmpMax > tmpMin {
		tmpIndex = (tmpMin + tmpMax) / 2
		if nums[tmpIndex] == target {
			return tmpIndex
		}
		if tmpIndex == tmpMin { //  已经找到了临界情况，即tmpMax = tmpMin+1
			if nums[tmpIndex] > target { //往数组左半部找
				return tmpMin
			} else {
				return tmpMax
			}
		}
		if nums[tmpIndex] > target { //往数组左半部找
			tmpMax = tmpIndex
		} else {
			tmpMin = tmpIndex
		}
	}
	return tmpIndex
}

// 成绩：双百