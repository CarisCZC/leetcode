package main

func replaceSpace(s string) string {
	res := ""
	offset := -1
	start := false
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if start {
				res += s[offset:i]
			}
			res += "%20"
			offset = i
			start = false
		} else if !start {
			offset = i
			start = true
		}
	}
	if start {
		res += s[offset:]
	}
	return res
}
