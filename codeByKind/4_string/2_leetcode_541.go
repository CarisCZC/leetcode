package main

// 反转字符串2
func reverseStr(s string, k int) string {
	res := []byte(s)
	for i := 0; i < len(res); {
		l := i
		j := min(l+k-1, len(res)-1)
		for l < j {
			res[l], res[j] = res[j], res[l]
			l++
			j--
		}
		i += 2 * k
	}
	return string(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	reverseStr("abcdefgdaiwnfgiadhadfnawifnah", 2)
}
