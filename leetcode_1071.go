// leetcode 1071 字符串的最大公因子

package main

import (
	"strings"
)

// 1. 如果str1 = nx，str2 = mx 那么str1+str2 = str2+str1 = (n+m)x
// func gcdOfStrings(str1 string, str2 string) string {
// 	if str1 +str2 == str2 +str1{
// 		return min(str1,str2)
// 	}
// 	return ""
// }

// 2. 根据前缀来匹配，也就是最长的前缀
// func gcdOfStrings(str1 string, str2 string) string {

// 	for i:=len(min(str1,str2));i>0;i--{
// 		if len(str1)%i==0 && len(str2)%i==0{
// 			tmp := str1[:i]
// 			if check(tmp,str1)&&check(tmp,str2){
// 				return tmp
// 			}

// 		}
// 	}
// 	return ""
// }

// 3. 优化的前缀匹配
// 有效的前缀串，其长度一定是str1和str2的最大公因子
func gcdOfStrings(str1,str2 string) string{
	lengcd := gcd(len(str1),len(str2))
	tmp := str1[:lengcd]
	if check(tmp,str1) && check(tmp,str2){
		return tmp
	}
	return ""
}

func gcd (a,b int) int{
	reminder := a%b
	for reminder !=0{
		a =b 
		b = reminder
		reminder = a%b
	}
	return b
}
func min(str1,str2 string) string{
	if len(str1) < len(str2){
		return str1
	}
	return str2
}

func check(tmp,str string) bool{
	lenx :=len(str)/len(tmp)
	ans := strings.Builder{}
	for i:=0;i<lenx;i++{
		ans.WriteString(tmp)
	}
	return str == ans.String()
}
