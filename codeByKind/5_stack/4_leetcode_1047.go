// https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/
package main

func removeDuplicates(s string) string {
	res := []byte{}
	offset := 0
	for i := 0; i < len(s); i++ {
		if offset > 0 {
			if res[offset-1] == s[i] {
				res = res[:offset-1]
				offset--
				continue
			}
		}
		res = append(res, s[i])
		offset++
	}
	return string(res)
}
