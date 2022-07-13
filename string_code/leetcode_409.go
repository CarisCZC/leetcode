//leetcode 409,最长回文串

package main

//判断字符是否出现过偶数次，最长回文串就是偶数次次数的和+1

func longestPalindrome(s string) int {
	tmpmap := map[rune]int{}

	for _, v := range s {
		tmpmap[v]++
	}
	res := 0
	flag := false
	for _, v := range tmpmap {
		if v%2 == 0 {
			res += v
		} else { //v是奇数
			// 奇数出现的次数 1，3，5，7，9
			if !flag {
				res += v
				flag = true
			} else {
				res += v - 1 //v-1是偶数
			}
		}
	}
	return res
}
