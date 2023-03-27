// https://leetcode.cn/problems/restore-ip-addresses/

package main

import (
	"strconv"
	"strings"
)

// 回溯分割
func restoreIpAddresses(s string) []string {
	path := []byte{}
	res := []string{}
	var backtrack func(s string, index int)
	dotNum := 0
	backtrack = func(s string, index int) {
		if index == len(s) && dotNum == 3 {
			res = append(res, string(path))
			return
		}
		if dotNum == 3 && index != len(s) {
			return
		}
		if index != 0 {
			dotNum++
			path = append(path, '.')
		}
		offset := 1
		for i := index; i < index+3; i++ {
			if i >= len(s) {
				break
			}
			// 先判断下合法性
			num, _ := strconv.Atoi(s[index : i+1])
			if num > 255 || (s[index] == '0' && i != index) { // 带有前导0
				break
			}
			path = append(path, s[index:i+1]...)
			backtrack(s, i+1)
			path = path[:len(path)+index-i-1]
			offset *= 10
		}
		if dotNum != 0 {
			dotNum--
			path = path[:len(path)-1]
		}

	}
	backtrack(s, 0)
	return res
}

// 优秀写法
func restoreIpAddresses2(s string) []string {
	path := []string{}
	res := []string{}
	var dfs func(s string, start int)
	dfs = func(s string, start int) {
		if len(path) == 4 {
			if start == len(s) {
				str := strings.Join(path, ".")
				res = append(res, str)
			}
			return
		}
		for i := start; i < len(s); i++ {
			if i != start && s[start] != 0 {
				break
			}
			str := s[start : i+1]
			num, _ := strconv.Atoi(str)
			if num >= 0 && num < 255 {
				path = append(path, str)
				dfs(s, i+1) //进入下一层
				path = path[:len(path)-1]
			} else {
				break
			}
		}
	}
	dfs(s, 0)
	return res
}
