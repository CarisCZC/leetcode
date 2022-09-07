// leetcode 599,两个列表的最小索引和
package main

// 思路，需要用一个map存储list1，K=v，V=idx；然后在list2中进行索引加和，然后用临时变量存储
func findRestaurant(list1 []string, list2 []string) []string {
	mp := make(map[string]int)
	for i, v := range list1 {
		mp[v] = i
	}
	res := make([]string, 1, len(list1))
	ans := 0
	flag := false
	for i, v := range list2 {
		if j, ok := mp[v]; ok {
			tmpans := i + j
			if !flag {
				ans += tmpans
				res[0] = v
				flag = true
			} else if tmpans < ans {
				res = res[:1]
				res[0] = v
				ans = tmpans
			} else if tmpans == ans {
				res = append(res, v)
			}
		}
	}

	return res
}

// 官方题解同等思路，然后有些奇特用法
// 	indexSum := math.MaxInt32
// 可以直接得出int32位的最大值
