package main

import "strconv"

// 分治思想，对于每个运算符号，分为左边结果和右边结果
// 于其说是动态规划，不如说是回溯
func diffWaysToCompute(expression string) []int {
	memo := map[string][]int{}
	return compute(expression, memo)
}
func compute(expression string, memo map[string][]int) []int {
	if ok, val := isNumber(expression); ok {
		return []int{val}
	}
	if val, ok := memo[expression]; ok {
		return val
	}

	res := make([]int, 0)
	for i := 0; i < len(expression); i++ {
		ch := expression[i]
		if ch < '0' || ch > '9' {
			left := compute(expression[:i], memo)
			right := compute(expression[i+1:], memo)
			val := 0
			for _, l := range left {
				for _, r := range right {
					switch ch {
					case '+':
						val = l + r
					case '-':
						val = l - r
					case '*':
						val = l * r
					}
					res = append(res, val)
				}
			}
		}
	}
	memo[expression] = res
	return res
}

func isNumber(expression string) (bool, int) {
	val, err := strconv.Atoi(expression)
	if err != nil {
		return false, -1
	}

	return true, val
}
