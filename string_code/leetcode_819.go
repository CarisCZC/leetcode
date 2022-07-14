// leetcode 819 最常见的单词

package main

import (
	"math"
	"strings"
	"unicode"
)

// 先统计出单词
// 大写也换小写
func mostCommonWord(paragraph string, banned []string) string {
	pre := 0
	flag := false
	smap := map[string]int{}
	for _, v := range banned {
		smap[v] = math.MinInt
	}
	newpara := paragraph + " "
	for i, v := range newpara {
		if unicode.IsLetter(v) && !flag {
			pre = i
			flag = true
		}
		if !unicode.IsLetter(v) && flag { //该位置非单词，且，pre已经记录
			tmp := strings.Builder{}
			for _, v := range newpara[pre:i] {
				tmp.WriteRune(unicode.ToLower(v))
			}
			smap[tmp.String()]++
			flag = false
		}
	}
	res := ""
	cnt := 0
	for k, v := range smap {
		if v > cnt {
			res, cnt = k, v
		}
	}
	return res
}
