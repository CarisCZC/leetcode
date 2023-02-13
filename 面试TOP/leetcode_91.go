package main

func numDecodings(s string) int {
	n := len(s)
	a, b, c := 0, 1, 0 // fi-2.fi-1,fi
	for i := 1; i <= n; i++ {
		c = 0
		if s[i-1] != '0' { //如果前一位为0，则当前为只能单独解析
			c += b
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			c += a
		}
		a, b = b, c
	}
	return c
}
