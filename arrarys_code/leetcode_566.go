// leetcode 566 ,重塑矩阵
package main

// 根据输入改造矩阵
func matrixReshape(mat [][]int, r int, c int) [][]int {
	nm := len(mat) * len(mat[0])
	if nm != r*c {
		return mat
	}
	ans := make([][]int, r)
	flag := 0
	for _, row := range mat {
		for _, v := range row {
			idx := flag / c
			if ans[idx] == nil {
				ans[idx] = make([]int, c)
			}
			ans[idx][flag%c] = v
			flag++
		}
	}
	return ans
}
