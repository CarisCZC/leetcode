//leetcode 13, 罗马数字转整数
package main

var temmap = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) int {
	// temmap := make(map[byte]int)
	// temmap['I'] = 1
	// temmap['V'] = 5
	// temmap['X'] = 10
	// temmap['L'] = 50
	// temmap['C'] = 100
	// temmap['D'] = 500
	// temmap['M'] = 1000
	sum := 0
	for i := range s {
		value := temmap[s[i]]
		if i < len(s)-1 && value < temmap[s[i+1]] {
			sum -= value
		} else {
			sum += value
		}
	}
	return sum
}
