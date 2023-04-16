package main

// 应该是要最大代价
// 满足的是一个二分搜索查找树
// 求个每个叶子节点的路径和，并取最大值
// f[i][j] 表示[i,j]范围内猜数字的最大代价
// f[i][j] = min{k+max(f(i,k-1),f(k+1,j))}
func getMoneyAmount(n int) int {
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			f[i][j] = j + f[i][j-1] // [i,j]的花费=[i,j-1]的花费+自身值
			for k := i; k < j; k++ {
				// 保证胜利的花费
				cost := k + max(f[i][k-1], f[k+1][j])
				if cost < f[i][j] { // 在所有保证胜利的花费中，找个最少的
					f[i][j] = cost
				}
			}
		}
	}
	return f[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
