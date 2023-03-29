package main

import "sort"

func findItinerary(tickets [][]string) []string {
	res := []string{}
	ticketMap := map[string][]string{}
	for _, ticket := range tickets {
		ticketMap[ticket[0]] = append(ticketMap[ticket[0]], ticket[1])
	}
	usedMap := map[string][]bool{}
	for start, arrive := range ticketMap {
		sort.Strings(ticketMap[start])
		usedMap[start] = make([]bool, len(arrive))
	}

	var backtarck func(arrive []string, start string)
	// 需要一个used数组，表示已经用了这个机票
	backtarck = func(arrive []string, start string) {
		if len(res) == len(tickets)+1 {
			return
		}
		if len(arrive) == 0 {
			return
		}

		for i := 0; i < len(arrive); i++ {
			if usedMap[start][i] {
				continue
			}
			res = append(res, arrive[i])
			usedMap[start][i] = true
			backtarck(ticketMap[arrive[i]], arrive[i])
			//判断，是否满足
			if len(res) == len(tickets)+1 {
				return
			}
			// 否则，说明不满足，就回退
			usedMap[start][i] = false
			res = res[:len(res)-1]
		}
	}
	res = append(res, "JFK")
	backtarck(ticketMap["JFK"], "JFK")
	return res
}

func main() {
	findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}})
}
