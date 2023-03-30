package main

// 电话号码字母组合
func letterCombinations(digits string) []string {
	res := []string{}
	cur := make([]byte, 0, len(digits))
	letterMap := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var backtrack func(digits string, start int)
	backtrack = func(digits string, start int) {
		if len(cur) == len(digits) {
			res = append(res, string(cur))
			return
		}
		digit := int(digits[start] - '0')
		str := letterMap[digit-2]
		for i := 0; i < len(str); i++ {
			cur = append(cur, str[i])
			backtrack(digits, start+1)
			cur = cur[:len(cur)-1]
		}
	}
	if digits == "" {
		return res
	}
	backtrack(digits, 0)
	return res
}
