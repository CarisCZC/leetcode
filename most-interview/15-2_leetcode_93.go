// https://leetcode.cn/problems/restore-ip-addresses/
// 93. 复原 IP 地址

package main

import "strconv"

// 这题回溯
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
		for i := index; i < index+3; i++ {
			if i >= len(s) {
				break
			}
			num, _ := strconv.Atoi(s[index : i+1])
			if num > 255 || (s[index] == '0' && i != index) {
				break
			}
			path = append(path, s[index:i+1]...)
			backtrack(s, i+1)
			path = path[:len(path)+index-i-1]
		}
		if dotNum != 0 {
			dotNum--
			path = path[:len(path)-1]
		}
	}

	backtrack(s, 0)
	return res
}
