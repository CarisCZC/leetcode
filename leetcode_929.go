// leetcode 929, 独特的邮件地址

package main

import "strings"

//  @之前的.不要，+之后的到@之前的内容不要
func numUniqueEmails(emails []string) int {
	tmp := map[string]bool{}
	for _, v := range emails {
		valid := strings.Builder{}
		for i, r := range v {
			if r == '+' { //就从后往前定位要的部分
				j := len(v) - 1
				for ; j > i; j-- {
					if v[j] == '@' {
						break
					}
				}
				valid.WriteString(v[j:])
				break
			}
			if r != '.' {
				if r == '@' {
					valid.WriteString(v[i:])
					break
				}
				valid.WriteRune(r)
			}
		}
		tmp[valid.String()] = true
	}
	return len(tmp)

}
