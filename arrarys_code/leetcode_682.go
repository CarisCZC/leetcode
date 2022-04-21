// leetcode 682, 棒球比赛

package main

import "strconv"

func calPoints(ops []string) int {
	cords := make([]int, 0, len(ops))
	for _, v := range ops {
		n := len(cords)
		if v == "+" {
			cords = append(cords, cords[n-1]+cords[n-2])
		} else if v == "C" {
			cords = cords[:n-1]
		} else if v == "D" {
			cords = append(cords, 2*cords[n-1])
		} else {
			num, _ := strconv.Atoi(v)
			cords = append(cords, num)
		}
	}
	ans := 0
	for _, v := range cords {
		ans += v
	}
	return ans
}
