package main

// 螺旋矩阵，模拟
// 保持区间一致性
func generateMatrix(n int) [][]int {
	res := make([][]int, 0, n)

	for i := 0; i < n; i++ {
		line := make([]int, n, n)
		res = append(res, line)
	}
	cnt := 1
	startx, starty := 0, 0
	mid := n / 2
	loop := n / 2
	offset := 1

	for loop > 0 {
		loop--
		i := startx
		j := starty
		// 从左到右
		for j = starty; j < (n - offset); j++ {
			res[startx][j] = cnt
			cnt++
		}
		// 右列从上到下
		for i = startx; i < (n - offset); i++ {
			res[i][j] = cnt
			cnt++
		}
		//下行从右到左
		for ; j > starty; j-- {
			res[i][j] = cnt
			cnt++
		}
		//左列从下到上
		for ; i > startx; i-- {
			res[i][j] = cnt
			cnt++
		}
		// 第二圈，起始位置改变
		startx++
		starty++
		// 遍历长度也减少1
		offset++
	}

	// 如果是奇数，中心位置需要处理
	if n%2 != 0 {
		res[mid][mid] = cnt
	}
	return res
}
