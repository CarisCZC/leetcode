// leetcode 661,图片平滑器
package main

//根据题的提示，其实可以看出相关的方法

// 首先，对于每个顶点来说，计算过程是一样的，然后对于左边中间部分的点，计算方法都一致
// 同理，上边，下边，右边，也是一致的。
// 官方写法，自己写的很复杂
func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			sum, num := 0, 0
			// 这里处理边界很巧妙
			for _, row := range img[max(i-1, 0):min(i+2, m)] {
				// 这里处理边界很巧妙
				for _, v := range row[max(j-1, 0):min(j+2, n)] {
					sum += v
					num++
				}
			}
			ans[i][j] = sum / num
		}
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
