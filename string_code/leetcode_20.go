// leetcode 20 ,正确的括号
package main

// 括号必须时正确的顺序
// ({)}不行，({})可以
func isValid(s string) bool {
	// 设置一个栈，可以以slice表达
	stack := make([]rune, 0, len(s))
	// stack = append(stack, rune(s[0]))
	smap := map[rune]rune{'}': '{', ']': '[', ')': '('}
	n := len(stack)
	for _, value := range s {
		if n == 0 {
			stack = append(stack, value)
			n++
			continue
		}
		if stack[n-1] == smap[value] {
			n--
			stack = stack[:n]
		} else {
			stack = append(stack, value)
			n++
		}
	}

	return n == 0

}

func main() {
	s := "()[]{}"
	isValid(s)
}

// 官方解答
// func isValid(s string) bool {
// 	n := len(s)
// 	if n%2 == 1 {
// 		return false
// 	}
// 	pairs := map[byte]byte{
// 		')': '(',
// 		']': '[',
// 		'}': '{',
// 	}
// 	stack := []byte{}
// 	for i := 0; i < n; i++ {
// 		if pairs[s[i]] > 0 { //paris中只有右括号，当查找左括号时，会查不到，返回0.因此我们要把左括号存进去
// 			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
// 				return false
// 			}
// 			stack = stack[:len(stack)-1]
// 		} else {
// 			stack = append(stack, s[i])
// 		}
// 	}
// 	return len(stack) == 0
// }
