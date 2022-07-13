// leetcode 125,验证回文串
// “上海自来水来自海上”，这就是回文串

package main

// 先处理字符串：非字母和数字字符去除
// 大写变成小写
// ASCII编码
// 0-9 48~57
// A-Z 65-90
// a-z 97-122
// 可以优化
// func isPalindrome(s string) bool {
// 	// 空字符串是回文串
// 	res := ""
// 	if s == "" {
// 		return true
// 	}
// 	for _, v := range s {
// 		intv := int(v)
// 		if intv >= 65 && intv <= 90 { //是大写字母
// 			intv += 32
// 		}
// 		if (intv >= 48 && intv <= 57) || (intv >= 97 && intv <= 122) {
// 			//要么是数字，要么是小写字母
// 			res += string(intv)
// 		}
// 	}
// 	n := len(res)
// 	for i := 0; i < n/2; i++ {
// 		if res[i] != res[n-i-1] {
// 			return false
// 		}
// 	}
// 	return true
// }

//优化版本
// 首部第一个字母或数字和尾部第一个字母或数字进行比对
func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
next:
	for i < j {
		// 如果当前位置不是字母或者数字
		a := int(s[i])
		if a >= 65 && a <= 90 { //是大写字母
			a += 32
		}
		for !((a >= 48 && a <= 57) || (a >= 97 && a <= 122)) {
			i++
			continue next
		}
		b := int(s[j])
		if b >= 65 && b <= 90 { //是大写字母
			b += 32
		}
		for !((b >= 48 && b <= 57) || (b >= 97 && b <= 122)) {
			j--
			continue next
		}
		if a != b {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

func main() {
	s := "Marge, let's \"[went].\" I await {news} telegram."
	isPalindrome(s)
}
