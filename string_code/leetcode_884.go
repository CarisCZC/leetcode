// leetcode 884 两句话中的不常见的单词

package main

// 在一个句子中，恰好只出现一次
// 先分词
func uncommonFromSentences(s1 string, s2 string) []string {
	space := -1
	tmp := map[string]int{}
	res := []string{}
	for i := 0; i < len(s1); i++ {
		if s1[i] == ' ' {
			tmp[s1[space+1:i]]++
			space = i
		}
	}
	//处理最后一个
	tmp[s1[space+1:]]++
	space = -1
	for i := 0; i < len(s2); i++ {
		if s2[i] == ' ' {
			tmp[s2[space+1:i]]++
			space = i
		}
	}
	tmp[s2[space+1:]]++

	for k, v := range tmp {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
