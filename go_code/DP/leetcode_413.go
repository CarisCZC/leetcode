package main

// 等差数列划分
// dp[i][j]表示从i到j构成一个等差数列
// dp[i][j+1] = dp[i][j]&& nums[j+1]-nums[j] == nums[j]-nums[j-1]
// dp[i+1][j] = dp[i][j] // 即i~j构成等差数列，i+1~j也是等差数列
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	res := 0
	if nums[1]-nums[0] == nums[2]-nums[1] {
		dp[0][2] = true
		res++
	}
	// 初始化从0开始的时候
	for i := 3; i < n; i++ {
		dp[0][i] = dp[0][i-1] && nums[i]-nums[i-1] == nums[i-1]-nums[i-2]
		if dp[0][i] {
			res++
		}
	}
	for i := 1; i < n-2; i++ {
		// 两个数不算等差数列
		dp[i][i+2] = nums[i+2]-nums[i+1] == nums[i+1]-nums[i]
		if dp[i][i+2] {
			res++
		}
		// 内部循环要从4个数开始
		for j := i + 3; j < n; j++ {
			if dp[i-1][j] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1] && nums[j]-nums[j-1] == nums[j-1]-nums[j-2]
			}
			if dp[i][j] {
				res++
			}
		}
	}
	return res
}

// 官方写法，O(n)，O(1)
func numberOfArithmeticSlices2(nums []int) (ans int) {
	n := len(nums)
	if n == 1 {
		return
	}
	// d表示两个数的公差，t表示当前数的贡献度
	d, t := nums[1]-nums[0], 0
	for i := 2; i < n; i++ {
		// 假设长度为k的等差数列有m个
		// 此时nums[k+1]-nums[k] =d
		// k+1 会在之前的基础上，
		// 新贡献长度为3，4，5，.....，k，k+1（3...k应该有m个，而k+1长度只会有一个） 共m+1个
		// 这时总数量=m+m+1
		if nums[i]-nums[i-1] == d {
			t++
		} else {
			d, t = nums[i-1]-nums[i], 0
		}
		ans += t
	}
	return
}
