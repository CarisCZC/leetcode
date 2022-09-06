//leetcode 121， 买卖股票的最佳时机
//其实就是找最小数后面的最大数

package main

func maxProfit(prices []int) (maxIncome int) {
	maxIncome = 0
	pricesLen := len(prices)
	tmpSell := prices[pricesLen-1]
	for i := pricesLen - 2; i >= 0; i-- {
		if tmpSell > prices[i] {
			tmpIncome := tmpSell - prices[i]
			if tmpIncome > maxIncome {
				maxIncome = tmpIncome
			}
		} else {
			tmpSell = prices[i]
		}
	}
	return
}
