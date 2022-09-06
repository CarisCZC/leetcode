// leetcode 11,盛最多的水，mid题

package main

// 从题来看，这是一个求面积的题
// 因此采用双指针，从两端往中间找
// 难点：在于思考怎么去移动双指针
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	res := -1
	for i < j {
		min := -1
		if height[i] > height[j] {
			min = height[j]
			j--
		} else {
			min = height[i]
			i++
		}
		// 这里i，j已经移动了，因此距离要补1
		tmpres := (j - i + 1) * min
		if tmpres > res {
			res = tmpres
		}
	}
	return res
}
