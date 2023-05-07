package main

// 如果[i,j]在s中出现,且s[k] = base[k]，那么如果s[j+1]=base[k+1]则[i,j+1]也出现过
func findSubstringInWraproundString(s string) int {
	base := "abcdefghijklmnopqrstuvwxyz"
	n := 26
	// dp[i]表示以base[i]为结尾的最大子串长度
	dp := [26]int{}
	// 初始化
	k := 1
	// 获得的是第一个字符在base中的下标
	idx := int(s[0] - 'a')
	dp[idx] = k
	idx++ //希望找到的字符下标
	for i := 1; i < len(s); i++ {
		if s[i] == base[idx%n] {
			k++
		} else {
			// 此时没有找到期望的字符，重置当前累加和小标
			k = 1
			idx = int(s[i] - 'a')
		}
		if k > dp[idx%n] {
			dp[idx%n] = k
		}
		idx++
	}

	// 结果累加
	ans := 0
	for _, v := range dp {
		ans += v
	}
	return ans
}

func main() {
	findSubstringInWraproundString("cac")
}
