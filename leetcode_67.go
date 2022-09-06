// leetcode 67,二进制求和
// 给你两个二进制字符串，返回它们的和（用二进制表示）
package main

import "fmt"

// go语言不能直接用二进制表示数
func addBinary(a string, b string) string {
	res := make([]byte, max(len(a), len(b))+1)
	add := false
	n := len(res) - 1
	i := len(a) - 1
	j := len(b) - 1
	for i > -1 && j > -1 {
		if a[i] == '1' && b[j] == '1' { //说明要进位
			if add {
				res[n] = '1'
			} else {
				res[n] = '0'
			}
			add = true //进位标志位
		} else if a[i] == '1' || b[j] == '1' { // 此位置只有一个1
			if add {
				res[n] = '0'
			} else {
				res[n] = '1'
				add = false
			}
		} else { //位置上没有1
			if add {
				res[n] = '1'
				add = false
			} else {
				res[n] = '0'
			}
		}
		i--
		j--
		n--
	}
	//循环完最小的字符串已经加完，但是进位并没有处理,需要把产生的进位加到长的字符串上
	if i != -1 {
		for ; i > -1; i-- {
			if add {
				if a[i] == '1' {
					res[n] = '0'
				} else {
					res[n] = '1'
					add = false
				}
			} else {
				res[n] = a[i]
			}
			n--
		}
	}
	if j != -1 {
		for ; j > -1; j-- {
			if add {
				if b[j] == '1' {
					res[n] = '0'
				} else {
					res[n] = '1'
					add = false
				}
			} else {
				res[n] = b[j]
			}
			n--
		}
	}
	if add {
		res[0] = '1'
		return string(res)
	} else {
		return string(res[1:])
	}

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))
}
