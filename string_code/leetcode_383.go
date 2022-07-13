// leetcode 383,赎金信

package main

func canConstruct(ransomNote string, magazine string) bool {
	tmp := map[rune]int{}
	for _, v := range magazine {
		tmp[v]++
	}
	for _, v := range ransomNote {
		tmp[v]--
		if tmp[v] < 0 {
			return false
		}
	}
	return true
}
