// leetcode 506, 相对名次
package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	mp := make(map[int]int, 0)
	for idx, v := range score {
		mp[v] = idx
	}
	sort.Sort(sort.Reverse(sort.IntSlice(score)))
	ans := make([]string, len(score))
	for i := 0; i < len(score); i++ {
		if i == 0 {
			ans[mp[score[i]]] = "Gold Medal"
		} else if i == 1 {
			ans[mp[score[i]]] = "Silver Medal"
		} else if i == 2 {
			ans[mp[score[i]]] = "Bronze Medal"
		} else {
			ans[mp[score[i]]] = strconv.Itoa(i + 1)
		}
	}
	return ans
}

// 官方写法
// 特别注明：sort.Slice只能在go1.8以上的版本来能够使用

// var desc = [3]string{"Gold Medal", "Silver Medal", "Bronze Medal"}

// func findRelativeRanks(score []int) []string {
//     n := len(score)
//     type pair struct{ score, idx int }
//     arr := make([]pair, n)
//     for i, s := range score {
//         arr[i] = pair{s, i}
//     }
//     sort.Slice(arr, func(i, j int) bool { return arr[i].score > arr[j].score })

//     ans := make([]string, n)
//     for i, p := range arr {
//         if i < 3 {
//             ans[p.idx] = desc[i]
//         } else {
//             ans[p.idx] = strconv.Itoa(i + 1)
//         }
//     }
//     return ans
// }
