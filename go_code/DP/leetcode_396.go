package main

// f(k) =  f(k-1)+sum-n*nums[n-k]
func maxRotateFunction(nums []int) int {
	sum := 0
	n := len(nums)
	for _, v := range nums {
		sum += v
	}
	f := 0

	for j := 0; j < n; j++ {
		f += j * nums[j]
	}
	ans := f
	for i := n - 1; i > 0; i-- {
		// 这个循环中，唯一的变量是nums[i],[i]确实要从最大开始
		// n-k = i，当k=1时，i= n-1
		// 返回来，当i=n-1时，求的就是f(1)
		f = f + sum - n*nums[i]
		if f > ans {
			ans = f
		}
	}
	return ans
}
