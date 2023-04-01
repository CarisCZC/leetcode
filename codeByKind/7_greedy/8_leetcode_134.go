// https://leetcode.cn/problems/gas-station/

package main

// 贪心吧，DP不动
func canCompleteCircuit(gas []int, cost []int) int {
	// DP表示当前站能走的最远的站。停在这的原因是因为该站无法到达下一站
	for i, n := 0, len(gas); i < n; {
		sumOfGas, sumOfCost, dp := 0, 0, 0
		for dp < n {
			j := (i + dp) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfCost > sumOfGas { // 站j已经走不到了
				break
			}
			dp++ // 继续往前走
		}
		if dp == n {
			return i
		} else {
			i += dp + 1
		}
	}
	return -1
}

func canCompleteCircuit2(gas []int, cost []int) int {
	// 汽车一定是从一个剩余汽油>0的站开始的
	// 走一圈的汽油剩余量 Allgas - Allcost，且>=0
	curRest, totalRest := 0, 0
	start := 0
	// 先从0开始跑，站i剩余汽油不够，起始位置变成i+1
	for i := 0; i < len(gas); i++ {
		curRest += gas[i] - cost[i] // 当前累加量
		totalRest += gas[i] - cost[i]
		if curRest < 0 { //只能从下一站出发
			start = i + 1
			curRest = 0 // 重新累加
		}
	}
	if totalRest < 0 {
		return -1
	}
	return start
}
