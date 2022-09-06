// leetcode 62 ,不同路径

package main

import "math/big"

func uniquePaths(m int, n int) int {

	// 必须使用big.Int
	return int(new(big.Int).Binomial(int64(m+n-2),int64(min(m,n))).Int64())
// 	res := uint64(1)
// 	a := uint64(m+n-2)
// 	b := uint64(min(m,n)-1)
// 	cut := uint64(1)
// 	for i:=0;i<min(m,n)-1;i++{
// 		res *= a
// 		cut *= b
// 		a--
// 		b--
// 	}
// 	return int(res/cut)
}

func min(a,b int)int{
	if a<= b{
		return a
	}
	return b
}