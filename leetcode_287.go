// leetcode 387,字符串中的第一个唯一字符

package main

// 创建一个26大小的数组，初始都为-1
// int[s[i]-'a'] = i,重复出现的话，后面的i会覆盖前面的i
type flag struct {
	Val   rune //值
	Index int  //出现的最小坐标
	Uni   int  //出现次数
}

func firstUniqChar(s string) int {
	a := [26]flag{}
	for i, v := range s {
		index := v - 'a'
		if a[index].Uni == 0 {
			a[index].Val = v
			a[index].Index = i

		}
		a[index].Uni++
	}
	index := -1
	for _, v := range a {
		if v.Uni == 1 {
			if index == -1 || v.Index < index {
				index = v.Index
			}
		}
	}
	return index
}
