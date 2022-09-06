// leetcode 1002 查找共用字符

package main

// 用一个map，表示某个字符出现的次数
func commonChars(words []string) []string {
	tmp := map[byte]int{}
	// 第一个word中的各个字符的出现次数
	for i:=0;i<len(words[0]);i++{
		tmp[words[0][i]]++
	}
	for i:=1;i<len(words);i++{
		tmp2 := map[byte]int{}
		for j:=0;j<len(words[i]);j++{
			if _,ok:=tmp[words[i][j]];ok{
				// 说明找到了
				tmp2[words[i][j]]++
			}
		}
		for k,v :=range tmp2{
			tmp2[k] = min(tmp[k],v)
		}
		tmp = tmp2
	}
	//根据tmp输出
	res := []string{}
	for k,v := range tmp{
		for i:=0;i<v;i++{
			res = append(res, string(k))
		}
	}
	return res
}
func min(a,b int)int{
	if a<=b{
		return a
	}else{
		return b
	}
}