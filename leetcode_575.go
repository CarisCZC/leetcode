// leetcode 575, 分糖果

package main

//统计一下有几种类型的糖果，如果种类数x>n/2,则可以吃到n/2种糖果，否则，只能吃到x种类
func distributeCandies(candyType []int) int {
	n := len(candyType) / 2
	mp := make(map[int]bool)
	for _, v := range candyType {
		if n == 0 {
			break
		}
		if !mp[v] {
			n--
			mp[v] = true
		}
	}
	return len(candyType)/2 - n
}
