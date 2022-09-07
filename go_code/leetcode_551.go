// leetcode 551, 学生出勤记录1

//
package main

func checkRecord(s string) bool {
	pcnt := 0
	Lcnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			pcnt++
			if pcnt >= 2 {
				return false
			}
		}
		if i <= len(s)-3 && s[i] == 'L' {
			Lcnt++
			for j := i + 1; j < i+3; j++ {
				if s[j] == 'L' {
					Lcnt++
				}
			}
			if Lcnt == 3 {
				return false
			} else {
				Lcnt = 0
			}
		}
	}
	return true
}
