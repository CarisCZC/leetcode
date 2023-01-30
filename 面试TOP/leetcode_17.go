// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/?favorite=2ckc81c
package main

func letterCombinations(digits string) []string {
	// 看用到了那几个数字，最多输入4个
	// useLetter := make([][]string, 0, 4)
	res := make([]string, 0, 128)
	letterMap := map[byte][]string{'2': {"a", "b", "c"}, '3': {"d", "e", "f"}, '4': {"g", "h", "i"}, '5': {"j", "k", "l"}, '6': {"m", "n", "o"}, '7': {"p", "q", "r", "s"}, '8': {"t", "u", "v"}, '9': {"w", "x", "y", "z"}}
	// 所有的都存入了useLetter
	res = combinationsLetter(digits, res, letterMap, "")
	return res
}

func combinationsLetter(digits string, res []string, letterMap map[byte][]string, tmp string) []string {
	if len(digits) == 0 {
		return res
	}
	letters := letterMap[digits[0]]
	if len(digits) == 1 {
		for _, ch := range letters {
			tmp += ch
			res = append(res, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	} else {
		for _, ch := range letters {
			res = combinationsLetter(digits[1:], res, letterMap, tmp+ch)
		}
	}
	return res
}

func main() {
	letterCombinations("234")
}
