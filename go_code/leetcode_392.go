// 392. 判断子序列
package main

// 用s里的在t中匹配，注意相对位置顺序即可
// 先比较s[0],确认位置t[a],然后以a为限位，确认s[1]，然后以t[b]
func isSubsequence(s string, t string) bool {
	n := 0
	for i := 0; i < len(t); i++ {
		if n >= len(s) {
			return true
		}
		if t[i] == s[n] {
			n++
		}
	}
	return n == len(s)
}
