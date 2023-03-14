package main

// he 242题的思想类似，主要的区别在于，要保存字符出现的最小次数
func commonChars(words []string) []string {
	wordHash := make([]int, 26, 26)
	for i, word := range words {
		if i == 0 {
			for _, ch := range word {
				wordHash[ch-'a']++
			}
			continue
		}
		otherWordHash := make([]int, 26, 26)
		for _, ch := range word {
			otherWordHash[ch-'a']++
		}
		for i := 0; i < 26; i++ {
			wordHash[i] = min(wordHash[i], otherWordHash[i])
		}
	}
	res := []string{}
	for ch, v := range wordHash {
		for i := 0; i < v; i++ {
			res = append(res, string(byte(ch+'a')))
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
