// leetcode 888 公平的糖果交换

package main

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	aliceMap := map[int]bool{}
	sumA,sumB :=0,0
	res := []int{}
	for i:=0;i<len(aliceSizes)||i<len(bobSizes);{
		if i<len(aliceSizes){
			sumA+=aliceSizes[i]
			aliceMap[aliceSizes[i]] = true
		}
		if i<len(bobSizes){
			sumB+=bobSizes[i]
		}
		i++
	}
	sumA = (sumA - sumB)/2
	for i:=0;i<len(bobSizes);i++{
		if aliceMap[bobSizes[i]+sumA]{
			res = []int{bobSizes[i]+sumA,bobSizes[i]}
			return res
		}
	}
return res
}
