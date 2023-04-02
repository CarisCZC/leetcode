package main

// 当碰到一个没有的字母时，就从此处从新划分。
// 如果后续又碰见了之前有的字母，那必须加入前面的字符段
func partitionLabels(s string) []int {
	res := []int{}
	letterMap := make(map[byte]int, 26) // 表示某个字母出现在哪一段内
	for i := 0; i < len(s); i++ {
		if k, ok := letterMap[s[i]]; ok { //如果出现过
			// 到i的所有单词，都必须放入res[k]里
			preIdx := 0
			for j := 0; j < len(res); j++ {
				if j <= k {
					preIdx += res[j]
				} else {
					res[k] += res[j]
				}
			}
			for j := preIdx; j < i; j++ { //改变后面单词出现的位置
				letterMap[s[j]] = k
			}
			res[k] += 1 // i加入进去
			res = res[:k+1]
		} else { // 该单词没有出现过
			res = append(res, 1)
			letterMap[s[i]] = len(res) - 1
		}
	}
	return res
}

func main() {
	partitionLabels("ababcbacadefegdehijhklij")
}
