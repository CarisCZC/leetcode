// leetcode 500, 键盘行

package main

func findWords(words []string) (ans []string) {
	line1 := "qwertyuiop"
	line2 := "asdfghjkl"
	line3 := "zxcvbnm"
	lineMap := make(map[rune]int)
	ans = make([]string, 0, len(words))
	for _, v := range line1 {
		lineMap[v] = 1
	}
	for _, V := range line2 {
		lineMap[V] = 2
	}
	for _, v := range line3 {
		lineMap[v] = 3
	}
	for _, v := range words {
		firstword := 0
		flag := false
		for i, word := range v {
			if word < 97 {
				word += 32
			}
			if firstword == 0 {
				firstword = lineMap[rune(word)]
			}
			if lineMap[rune(word)] != firstword {
				break
			}
			if i == len(v)-1 {
				flag = true
			}
		}
		if flag {
			ans = append(ans, v)
		}
	}
	return ans
}

// func main() {
// 	words := []string{"Hello", "Alaska", "Dad", "Peace"}
// 	fmt.Println(findWords(words))
// 	// fmt.println(findWords(words))
// }

// 官方写法，比较自己的写法在空间上是更优的
// 对于next的用法也是很巧妙的，这样就不需要使用flag
// func findWords(words []string) (ans []string) {
//     const rowIdx = "12210111011122000010020202"
// next:
//     for _, word := range words {
//         idx := rowIdx[unicode.ToLower(rune(word[0]))-'a']
//         for _, ch := range word[1:] {
//             if rowIdx[unicode.ToLower(ch)-'a'] != idx {
//                 continue next
//             }
//         }
//         ans = append(ans, word)
//     }
//     return
// }
