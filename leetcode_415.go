// leetcode 415 字符串相加

// 输入：num1 = "11", num2 = "123"
// 输出："134"
package main

import "strconv"

// 和之前的二进制相加是一样的
func addStrings(num1 string, num2 string) string {
	a := len(num1)
	b := len(num2)
	n := max(a, b)
	flag := 0
	res := ""
	a--
	b--
	for i := n; i >= 0; i-- {
		tmp := 0
		if a >= 0 && b >= 0 {
			tmp = int(num1[a] - '0' + num2[b] - '0')
			// res += strconv.Itoa((flag + tmp) % 10)
			// flag = (flag + tmp) / 10
			a--
			b--
		} else if b >= 0 { //说明a已经到头了
			tmp = int(num2[b] - '0')
			// res += strconv.Itoa((flag + tmp) % 10)
			// flag = (flag + tmp) / 10
			b--
		} else if a >= 0 { //说明b已经到头了
			tmp = int(num1[a] - '0')
			// res += strconv.Itoa((flag + tmp) % 10)
			// flag = (flag + tmp) / 10
			a--
		} else { // a,b都没了
			if flag == 1 {
				res = strconv.Itoa(flag) + res
			}
			return res
		}
		res = strconv.Itoa((tmp+flag)%10) + res
		flag = (flag + tmp) / 10
	}
	return res
}

// 代码更加简洁，但没什么本质区别
// func addStrings2(num1, num2 string) string {
// 	add := 0
// 	ans := ""
// 	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
// 		var x, y int
// 		if i >= 0 {
// 			x = int(num1[i] - '0')
// 		}
// 		if y >= 0 {
// 			y = int(num2[j] - '0')
// 		}
// 		res := x + y + add
// 		ans = strconv.Itoa(res%10) + ans
// 		add = res / 10
// 	}
// 	return ans
// }

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	num1 := "11"
	num2 := "123"
	addStrings(num1, num2)
}
