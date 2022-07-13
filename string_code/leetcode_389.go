// leetcode 389,找不同
package main

// 根据s创建一个map
// 然后根据t遍历这个map
// func findTheDifference(s string, t string) byte {
// 	tmp := map[rune]int{}
// 	for _, v := range s {
// 		tmp[v]++

// 	}
// 	res := ' '
// 	for _, v := range t {
// 		tmp[v]--
// 	}
// 	for k, v := range tmp {
// 		if v == 1 {
// 			res = k
// 		}
// 	}
// 	return byte(res)
// }

//官方解法：
//如果将两个字符串拼接成一个字符串，则问题转换成求字符串中出现奇数次的字符。
//类似于「136. 只出现一次的数字」，我们使用位运算的技巧解决本题。

func findTheDifference(s string, t string) (diff byte) {
	for i := range s {
		diff ^= s[i] ^ t[i]
	}
	return diff ^ t[len(t)-1]
}
