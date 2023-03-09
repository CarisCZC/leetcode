package main

func canCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		for cnt < n {
			// i表示出发的站，cnt表示从这个站能走的最远的站
			j := (i + cnt) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfCost > sumOfGas {
				// 说明第j站走不到
				break
			}
			// 能到，就继续走
			cnt++
		}
		if cnt == n {
			return i
		} else {
			// 在cnt内的站有不用考虑：因为从i到cnt+i不行，那么i+1到cnt+i也是不行的。
			i += cnt + 1
		}

	}
	return -1
}
