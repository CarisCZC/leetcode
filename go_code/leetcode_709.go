// leetcode 709,转换成小写字母

package main

func toLowerCase(s string) string {
	res := []byte(s)
	for i := 0; i < len(res); i++ {
		if res[i] >= 65 && res[i] <= 90 {
			res[i] += 32
		}
	}
	return string(res)
}
