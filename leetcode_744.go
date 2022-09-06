// leetcode 744,比目标字母大的最小字母

package main

func nextGreatestLetter(letters []byte, target byte) (res byte) {
	n := 0
	for ; n < len(letters); n++ {
		if target < letters[n] {
			return letters[n]
		}
	}
	return letters[n%len(letters)]
}
