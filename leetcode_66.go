// leetcode,66 加1
package main

func plusOne(digits []int) []int {
	tmpLen := len(digits)
	if digits[tmpLen-1] != 9 {
		digits[tmpLen-1] += 1
		return digits
	}
	for i := tmpLen - 1; i >= 0; i-- { //一旦进入循环，说明最后一位是9
		if digits[i] == 9 {
			if i == 0 {
				digits[i] = 1
				digits = append(digits, 0)
				return digits
			}
			digits[i] = 0
		} else {
			digits[i] += 1
			return digits
		}
	}
	return digits

}
