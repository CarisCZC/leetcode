// leetcode 953,验证外星语词典
package main

// func isAlienSorted(words []string, order string) bool {
// 	orderMap := map[byte]int{0:-1}//字符即相应顺序
// 	for i:=0;i<len(order);i++{
// 		orderMap[order[i]] = i
// 	}
// 	for i:=0;i<len(words);i++{
// 		next:
// 		for j:=i+1;j<len(words);j++{
// 			k1,k2:=0,0
// 			for ;k1<len(words[i])&&k2<len(words[j]);k1,k2=k1+1,k2+1{
// 				if words[i][k1] != words[j][k2]{
// 					if orderMap[words[i][k1]]>=orderMap[words[j][k2]]{
// 						return false
// 					}else{
// 						continue next
// 					}
// 				}
// 			}
// 			if k1<len(words[i]){ //循环因k2=len(words[j])结束，说明j比i短。
// 				return false;
// 			}
// 		}
// 	}
// 	return true;
// }

// 优秀解法
func isAlienSorted(words []string, order string) bool {
	bs := make([]int,26)
	for i:=0;i<26;i++{
		bs[order[i]-'a'] = i;
	}
	next:
	for i:=1;i<len(words);i++{
		for j:=0;j<len(words[i])&&j<len(words[i-1]);j++{
			pre,cur := bs[words[i-1][j]-'a'],bs[words[i][j]-'a'];
			if pre > cur {
				return false
			}
			if pre < cur{
				continue next
			}
		}
		if len(words[i-1]) > len(words[i]){
			return false
		}
	}
	return true
}