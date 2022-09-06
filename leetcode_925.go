// leetcode 925 长按键入

package main

// typed中的同字符忽略
func isLongPressedName(name string, typed string) bool {
	if len(typed) < len(name) {
		return false
	}
	for i, j := 0, 0; i < len(name) || j < len(typed); {
		if j == len(typed) && i != len(name) {
			return false
		}
		if i == len(name) {
			i--
		}
		if j == len(typed) {
			j--
		}
		if name[i] != typed[j] {
			//如果第一个字符都错误，那么一定是错的
			if i == 0 || j == 0 {
				return false
			}
			if typed[j] == typed[j-1] { //即是一个重复输入字符
				j++
				continue
			}
			return false
		}
		// 走到这说明字符相等
		i++
		j++
	}
	return true
}

// 优秀写法
// func isLongPressedName(name string, typed string) bool {
// 	i, j := 0, 0
// 	for j < len(typed) {
// 		if i < len(name) && name[i] == typed[j] {
// 			i++
// 			j++
// 		} else if j > 0 && typed[j] == typed[j-1] {
// 			j++
// 		} else {
// 			return false
// 		}
// 		//j++
// 	}
// 	return i == len(name)
// }
