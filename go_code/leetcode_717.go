// leetcode 717，1比特字符和2比特字符

package main

//如果bits的长度为2n，此时[100010]
// 最后一位一定是0，倒数第二位如果是0，那么最后一位一定会翻译成0
// 那么需要处理的是，倒数第二位是1的情况。那么倒数第三位也一定是1.
func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	sum := 0
	// 此时倒数第三位也一定是1
	for i := 0; i < n-1; {
		if bits[i] == 1 {
			sum += 2
			i += 2
		} else {
			sum += 1
			i++
		}
	}
	if sum == n-1 { //意味着前面的刚好解释OK
		return true
	} else {
		return false
	}
}
