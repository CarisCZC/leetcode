package main

func lemonadeChange(bills []int) bool {
	changeSum := make([]int, 3, 3) //5,10,20
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			changeSum[0]++
		}
		if bills[i] == 10 {
			changeSum[0]--
			if changeSum[0] < 0 {
				return false
			}
			changeSum[1]++
		}
		if bills[i] == 20 {
			if changeSum[1] > 0 {
				changeSum[1]--
			} else {
				changeSum[0] -= 2
			}
			changeSum[0]--
			if changeSum[0] < 0 || changeSum[1] < 0 {
				return false
			}
			changeSum[2]++
		}
	}
	return true
}
