package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return backtrack(1, n)
}

func backtrack(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTrees := []*TreeNode{}
	// 树一定会用到回溯
	// 这题的重点是怎么获得左右子树。
	for i := start; i <= end; i++ {
		leftTrees := backtrack(start, i-1)

		rightTrees := backtrack(i+1, end)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				curTree := &TreeNode{i, left, right}
				allTrees = append(allTrees, curTree)
			}
		}
	}
	return allTrees
}
