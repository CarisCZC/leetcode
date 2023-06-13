// https://leetcode.cn/problems/generate-parentheses/
// 22. 括号生成

package main

func generateParenthesis(n int) []string {
	pattern := make([]byte, n*2)
	pattern[0] = '('
	res := []string{}
	return generate(pattern, 1, 1, res)

}

// n 来表示当前括号的平衡情况，增加一个'(',n+1, 增加一个 )，n-1
// 当n==0时且index==len(pattern)。合法
func generate(pattern []byte, n, index int, res []string) []string {
	if n < 0 {
		return res
	}
	if index == len(pattern) {
		if n == 0 {
			res = append(res, string(pattern))
			return res
		}
		return res
	}
	pattern[index] = '('
	res = generate(pattern, n+1, index+1, res)
	pattern[index] = ')'
	res = generate(pattern, n-1, index+1, res)
	return res
}
