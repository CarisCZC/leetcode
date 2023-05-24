// https://leetcode.cn/problems/add-strings/
// 415. 字符串相加

package main

func addStrings(num1 string, num2 string) string {
	var res []byte
	i := len(num1) - 1
	j := len(num2) - 1
	up := '0'
	for i > -1 && j > -1 {
		cur := num1[i] - '0' + num2[j] + byte(up) - '0'
		up = '0'
		if cur > '9' {
			cur -= ('9' + 1 - '0')
			up += 1
		}
		res = append(res, cur)
		i--
		j--
	}
	for i > -1 {
		cur := num1[i] + byte(up) - '0'
		up = '0'
		if cur > '9' {
			cur -= ('9' + 1 - '0')
			up += 1
		}
		res = append(res, cur)
		i--
	}
	for j > -1 {
		cur := num2[j] + byte(up) - '0'
		up = '0'
		if cur > '9' {
			cur -= ('9' + 1 - '0')
			up += 1
		}
		res = append(res, cur)
		j--
	}
	if up > '0' {
		res = append(res, byte(up))
	}
	i = 0
	j = len(res) - 1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return string(res)
}
