// leetcode 748,最短补全词
package main

// 先将licensePlate中的字母提出来
func shortestCompletingWord(licensePlate string, words []string) (value string) {
	onlych := map[rune]int{}
	for _, v := range licensePlate {
		if v >= 'A' && v <= 'Z' {
			v += 32
		}
		if v >= 'a' && v <= 'z' {
			onlych[v]++
		}
	}
next:
	for _, v := range words {
		tmp := map[rune]int{}
		for _, k := range v {
			tmp[k]++
		}
		for k, va := range onlych {
			if tmp[k] < va {
				continue next
			}
		}
		if value == "" || len(v) < len(value) {
			value = v
		}
	}
	return
}
