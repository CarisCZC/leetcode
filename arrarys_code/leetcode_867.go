// leetcode 867,转置矩阵

package main


func transpose(matrix [][]int) [][]int {
	res := [][]int{}
	for i:=0;i<len(matrix[0]);i++{
		line := []int{}
		for j:=0;j<len(matrix);j++{
			line = append(line, matrix[j][i])
		}
		res = append(res, line)
	}
	return res
}     