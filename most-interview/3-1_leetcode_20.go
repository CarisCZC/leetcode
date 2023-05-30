// https://leetcode.cn/problems/valid-parentheses/
// 20. 有效的括号
package main

// 经典栈和map运用
func isValid(s string) bool {
	kuohao := map[byte]byte{'(': ')', '{': '}', '[': ']'}
	need := []byte{}
	id := -1
	for i := range s {
		if left, ok := kuohao[s[i]]; ok {
			need = append(need, left)
			id++
		} else if id != -1 && s[i] == need[id] { //不存在只能是右括号
			need = need[:id]
			id--
		} else {
			return false
		}
	}
	if id == -1 {
		return true
	}
	return false
}

func main() {
	isValid("(){}[]")
}
