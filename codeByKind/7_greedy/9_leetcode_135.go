package main

// 贪心，但感觉DP更好
func candy(ratings []int) int {
	// DP[i] 表示i个孩子获得的糖果。DP[0] = 1
	// ratings[i+1]>ratings[i],则DP[i+1]= DP[i]+1
	// ratings[i+1]<ratings[i],则DP[i+1]=DP[i],DP[i]++,
	// ratings[i]>ratings[i-1],则DP[i] = DP[i-1]+1
	// ratingsp[i]<ratings[i-1],则DP[i] = 1
	DP := make([]int, len(ratings), len(ratings))
	DP[0] = 1
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i+1] > ratings[i] {
			DP[i+1] = DP[i] + 1
		} else {
			DP[i+1] = 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			DP[i-1] = max(DP[i]+1, DP[i-1])
		}
	}
	res := 0
	for i := 0; i < len(DP); i++ {
		res += DP[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	candy([]int{1, 0, 2})
}
