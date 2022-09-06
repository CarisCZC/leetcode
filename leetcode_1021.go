// leetcode 1021 删除最外层的括号

package main

import "strings"

//记录每个闭合的场景
// 每一个(计数+1，每一个)计数减1，当计数等于0时，返回pre，i之间的string
func removeOuterParentheses(s string) string {
	pre := 1
	cnt := 1//第一个(默认在
	res := strings.Builder{}
	for i:=1;i<len(s);i++{
		if s[i] =='('{
			cnt++
		}else {
			cnt--
		}
		if cnt ==0{
			res.WriteString(s[pre:i])
			pre = i+2
		}
	}
	return res.String()
}