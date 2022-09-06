// leetcode 1078 Bigram分词

package main

import "strings"

func findOcurrences(text string, first string, second string) []string {
	texts := strings.Split(text, " ")
	res := []string{}
	for i:=0;i<len(texts)-2;i++{
		if texts[i] == first && texts[i+1] == second{
			res = append(res, texts[i+2])
		}
	}
	return res
}
func main(){
str :="alice is a good girl she is a good student"
first :="a"
second :="good"
findOcurrences(str,first,second)
}