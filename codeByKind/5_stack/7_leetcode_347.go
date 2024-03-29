// https://leetcode.cn/problems/top-k-frequent-elements/
package main

import (
	"container/heap"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	numMap := map[int]int{}
	for _, num := range nums {
		numMap[num]++
	}

	ans := []int{}
	for key, _ := range numMap {
		ans = append(ans, key)
	}
	// 构建k排序队列
	sort.Slice(ans, func(i, j int) bool {
		return numMap[ans[i]] > numMap[ans[j]]
	})
	return ans[:k]
}

// heap的用法
func topKFrequent2(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
