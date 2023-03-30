package main

// var (
// 	path []int
// 	res  [][]int
// )

func combine(n int, k int) [][]int {
	path := []int{}
	res := [][]int{}
	var backtrack func(n, k, index int)
	backtrack = func(n int, k int, index int) {
		if len(path) == k { //终止条件
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := index; i <= n; i++ {
			if n-i+1 < k-len(path) { // 即剩下的元素个数<k时，不用再循环回溯了,剪枝优化
				break
			}
			path = append(path, i)
			backtrack(n, k, i+1)
			path = path[:len(path)-1] // 回溯，刚刚的append
		}
	}
	backtrack(n, k, 1)
	return res
}

// func backtrack(n, k, start int) {
// 	if len(path) == k { //终止条件
// 		tmp := make([]int, k)
// 		copy(tmp, path)
// 		res = append(res, tmp)
// 		return
// 	}
// 	for i := start; i <= n; i++ {
// 		if n-i+1 < k-len(path) { // 即剩下的元素个数<k时，不用再循环回溯了
// 			break
// 		}
// 		path = append(path, i)
// 		backtrack(n, k, i+1)
// 		path = path[:len(path)-1] // 回溯，刚刚的append
// 	}

// }
