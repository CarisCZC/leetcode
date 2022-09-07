// leetcode 771,宝石与石头

package main

// map统计
func numJewelsInStones(jewels string, stones string) (ans int) {
	stonemap := map[rune]int{}
	for _, v := range stones {
		stonemap[v]++
	}
	for _, v := range jewels {
		ans += stonemap[v]
	}
	return
}
