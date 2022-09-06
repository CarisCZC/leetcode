// leetcode 860 柠檬水找0

package main

func lemonadeChange(bills []int) bool {
	charge := make([]int,3);
	// charge := 0;
	for _,v := range bills{
		index := v/10;
		charge[index] ++
		if index == 1{
			charge[index-1]--
			if charge[index-1]<0{
				return false
			}
		}
		if index == 2{
			charge[1]--
			charge[0]--
			if charge[1]<0{
				charge[1] = 0
				charge[0] -= 2
			}
			if charge[0]<0{
				return false
			}
		}
	}
	return true
}