// leetcode 228, 汇总区间
package main

import "strconv"

func summaryRanges(nums []int) (res []string) {
	if len(nums) == 0 {
		res = make([]string, 0)
		return
	}
	if len(nums) == 1 {
		res = make([]string, 1)
		res[0] = strconv.Itoa(nums[0])
		return
	}
	preNum := nums[0]
	tmpLen := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != tmpLen+preNum+1 { //也就是不连续
			if tmpLen == 0 {
				res = append(res, strconv.Itoa(preNum))
			} else {
				res = append(res, strconv.Itoa(preNum)+"->"+strconv.Itoa(nums[i-1]))
			}
			preNum = nums[i]
			tmpLen = 0
		} else {
			tmpLen++
		}
	}
	// 处理最后一个数据
	if tmpLen == 0 { //最后一个元素与前面的不连续
		res = append(res, strconv.Itoa(preNum))
	} else {
		res = append(res, strconv.Itoa(preNum)+"->"+strconv.Itoa(preNum+tmpLen))
	}
	return
}
