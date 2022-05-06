// leetcode 40 , 组合总和Ⅱ
// 要求所有的数字只能使用一次
package main

import (
	"fmt"
	"sort"
)

// 总体上说算法没有什么区别，上次的回溯这次依然可以用
func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	findCombination2(&res, []int{}, candidates, 0, target)
	return res

}

// 将tmpCom存入res时，有时会出现修改tmpCom时res内部的值也会修改。
func findCombination2(res *[][]int, tmpCom, candidates []int, tmpSum, target int) {
	for i, v := range candidates {
		if i > 0 && v == candidates[i-1] {
			continue
		}
		tmpSum += v
		tmpCom = append(tmpCom, v)
		if tmpSum == target {
			newCom := make([]int, len(tmpCom))
			copy(newCom, tmpCom)
			*res = append(*res, newCom)
			return
		}
		if tmpSum > target {
			return
		}
		findCombination2(res, tmpCom, candidates[i+1:], tmpSum, target)
		n := len(tmpCom)
		tmpSum -= v
		tmpCom = tmpCom[:n-1]
	}
}
func main() {
	nums := []int{3, 1, 3, 5, 1, 1}
	target := 8
	fmt.Print(combinationSum2(nums, target))
}
