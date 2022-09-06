// leetcode 896, 单调数列

package main


func isMonotonic(nums []int) bool {
	flag := 0
	for i := 0;i<len(nums)-1;i++{
		j := i+1
		if nums[i] != nums[j]{
			if flag == 0{
				flag = plus(nums[i],nums[j])
			}else{
				if flag != plus(nums[i],nums[j]){
					return false
				}
			}
		}
	}
	return true

}
// 根据前两个元素，确定递增还是递减
// true递增，false递减
func plus(a,b int) int{
	if a<b{
		return 1
	}
	return -1
}