// https://leetcode.cn/problems/generate-parentheses/?favorite=2ckc81c

package main

func generateParenthesis(n int) []string {
	char := make([]byte, 2*n, 2*n)
	char[0] = '('
	char[2*n-1] = ')'
	res := []string{}
	res = putLeft(res, 1, 1, char)
	return res
}

//
func putLeft(res []string, index, left int, char []byte) []string {
	n := len(char) / 2
	if left < 0 || left > n {
		return res
	}
	if index == len(char)-1 {
		if left == 1 {
			res = append(res, string(char))
		}
		return res
	}
	char[index] = '('
	res = putLeft(res, index+1, left+1, char)
	char[index] = ')'
	res = putLeft(res, index+1, left-1, char)
	return res
}
