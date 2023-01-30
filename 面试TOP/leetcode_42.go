// 接雨水

package main

func trap(height []int) int {
	leftIndex := 0
	rightIndex := 0
	res := 0
	for i := 1; i < len(height)-1; i++ {

		cur := height[i]
		if cur > height[leftIndex] {
			leftIndex = i
			continue
		}

		// 找右边最高
		if i > rightIndex {
			tmp := height[i+1]
			rightIndex = i + 1
			for right := i + 1; right < len(height); right++ {
				if height[right] > tmp {
					rightIndex = right
					tmp = height[right]
				}
			}
			// 右边最高没他高
			if height[rightIndex] < cur {
				continue
			}
		}
		res += min(height[leftIndex], height[rightIndex]) - cur
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
