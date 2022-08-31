// leetcode 1108,IP地址无效化

// 用[.]替换字符串中的每个.
package main

import "strings"

func defangIPaddr(address string) string {
	ans := strings.Builder{}
	for _,v := range address{
		if v == '.'{
			ans.WriteString("[.]")
		}else{
			ans.WriteRune(v)
		}
	}
	return ans.String()
}