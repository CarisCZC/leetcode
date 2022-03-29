// leetcode 26. 删除有序数组中的重复项
package main

// 遍历一遍数组，如果有不重复的，就进行移位
func removeDuplicates(nums []int) (len int) {
	len = 1	//nums[0]默认是不重复的
	for i,v:= range nums {
		tmpValue := nums[len-1]
		if v == tmpValue{ //即和前一个数字重复
			continue
		}else {
			len++
			nums[len-1] = nums[i] 
		}
	}
	return 

}
