//leetcode 136, 只出现一次的数字
// 对于常规写法，就是使用hashMap来进行查找
// 而更进一步的写法，是位运算方法
// 与运算 1&1 = 1, 1&0=0
// (相同位的两个数字都为1，则为1；若有一个不为1，则为0)
// 或运算 1|0=1，相同位只要有一个为1即为1
// 异或运算 1^0 =1 ,相异为1，相同则为0

package main

func singleNumber(nums []int) int {
	single := nums[0]
	for i := 1; i < len(nums); i++ {
		single ^= nums[i]
	}
	return single
}
