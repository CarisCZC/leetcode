// leetcode 303, 区域与检索——数组不可变

package main

type NumArray struct {
	nums    []int
	NumsLen int
}

// func Constructor(nums []int) (newArray NumArray) {
// 	newArray.nums = make([]int, len(nums))
// 	copy(newArray.nums, nums)
// 	newArray.NumsLen = len(nums)
// 	return
// }

// func (this *NumArray) SumRange(left int, right int) int {
// 	if left < 0 || left >= this.NumsLen || right >= this.NumsLen {
// 		panic("error left or right")
// 	}
// 	sum := 0
// 	for ; left <= right; left++ {
// 		sum += this.nums[left]
// 	}
// 	return sum

// }

// 前缀和 写法
// 题中隐藏了一个内容，即NumArray与nums不必一一对应，因此NumArray中的数值不一定时初始化中的nums值
func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i, v := range nums {
		sums[i+1] = sums[i] + v
	}
	return NumArray{sums, len(sums)}
}

// 返回特定位置的值
// func (this *NumArray) find(index int) int {
// 	return this.nums[index+1] - this.nums[index]
// }

func (this *NumArray) SumRange(left int, right int) int {
	return this.nums[right+1] - this.nums[left]

}
