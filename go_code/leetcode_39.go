// leetcode 39,组合总和

package main

import (
	"fmt"
	"sort"
)

//
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates) //先排个序
	findCombination(&res, []int{}, candidates, 0, target)
	return res
}

func findCombination(combination *[][]int, tmpCom, candidates []int, tmpSum, target int) { // combination用来存放结果,tmpSum用于存阶段值
	for i, v := range candidates {
		tmpCom = append(tmpCom, v)
		n := len(tmpCom)
		tmpSum += v
		if tmpSum == target {
			newCom := make([]int, n)
			copy(newCom, tmpCom)
			*combination = append(*combination, newCom) //将组合加入
			return                                      //这里就没有第二种可能了
		}
		if tmpSum > target { // 回溯tmpCom和tmpSum
			break
		}
		findCombination(combination, tmpCom, candidates[i:], tmpSum, target) //不用比自身小的数
		tmpSum -= v
		tmpCom = tmpCom[:n-1]
	}
}

func main() {
	nums := []int{2, 3, 5}
	target := 8
	fmt.Print(combinationSum(nums, target))
}
