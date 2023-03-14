package main

// 原地修改
func reverseWords(s string) string {
	//双指针去空格
	slow, fast := 0, 0
	words := []byte(s)
	for fast < len(words) && words[fast] == ' ' {
		fast++
	}
	// 删除单词中间的冗余空格
	for ; fast < len(words); fast++ {
		if fast-1 > 0 && words[fast-1] == words[fast] && words[fast] == ' ' {
			continue
		}
		words[slow] = words[fast]
		slow++
	}
	// slow是字符串的有效长度，超过slow的部分应全部删除
	// 删除后面的空格
	if words[slow-1] == ' ' {
		words = words[:slow-1]
	} else {
		words = words[:slow]
	}
	// 2. 反转整个字符串
	reverse(&words, 0, len(words)-1)
	// 3. 翻转每个单词
	for i := 0; i < len(words); {
		j := i
		for ; j < len(words) && words[j] != ' '; j++ {

		}
		reverse(&words, i, j-1)
		i = j
		i++
	}
	return string(words)
}

func reverse(s *[]byte, left, right int) {
	words := *s
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
}

func main() {
	reverseWords("  hello   world  ")
}
