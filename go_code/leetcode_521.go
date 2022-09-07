//leetcode 521,最长特殊序列1

//a不是b的子序列（b>a），那么，最长就是b
//a是b的子序列，那么，最长就是b

// 那么，问题就成了，判断a是不是b的子序列？

package main

func findLUSlength(a string, b string) int {
	a, b = maxlen(a, b)
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return len(b)
			}
		}
		return -1

	}
	return len(b)
}

func maxlen(a, b string) (string, string) {
	if len(a) >= len(b) {
		return b, a
	}
	return a, b
}
