// 最小操作次数使数组元素相等

// 原题说每次会使n-1个元素增加1，反向思考其实就是每次使一个数-1
// 那么每个数与最小值之间的差值，就是这个数需要被操作的次数
package main

func minMoves(nums []int) int {
	min := nums[0]
	moves := 0
	for i, v := range nums {
		if v >= min {
			moves += v - min
		} else {
			baseMove := min - v
			moves += i * baseMove
			min = v
		}
	}
	return moves
}
