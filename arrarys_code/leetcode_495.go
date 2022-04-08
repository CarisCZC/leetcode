// leetcode 495, 提莫攻击

// 可以理解为找一个最长的连续区间
// 给的数字可能已经在这个区间里了，因此有重复那种可能，要去除重复的区间
package main

// func findPoisonedDuration(timeSeries []int, duration int) int {
// 	ans := duration //第一个至少中毒这么长时间
// 	tmpAns := 0
// 	lastAns := 0
// 	for i := 1; i < len(timeSeries); i++ {
// 		lastAns = timeSeries[i-1] + duration - 1
// 		tmpAns = timeSeries[i] + duration - 1
// 		if timeSeries[i] > lastAns {
// 			ans += duration
// 		} else {
// 			ans += tmpAns - lastAns
// 		}
// 	}
// 	return ans
// }

// 算区间
// 从算法分析角度来说，这个更优，因为每次循环少访问了一次timeSeries[i-1]
// 实际上来说，二者差别不大，都是O(n)的算法
func findPoisonedDuration(timeSeries []int, duration int) int {
	ans := 0
	thisV := timeSeries[0]
	thisSeries := timeSeries[0] + duration - 1
	for i := 1; i < len(timeSeries); i++ {
		tmpSeries := timeSeries[i] + duration - 1
		if timeSeries[i] > thisSeries { // 说明断开，把之前的连续区间加上，然后重置起点
			ans += thisSeries - thisV + 1
			thisSeries = tmpSeries
			thisV = timeSeries[i]
		} else {
			thisSeries = tmpSeries
		}
	}
	ans += thisSeries - thisV + 1
	return ans
}
