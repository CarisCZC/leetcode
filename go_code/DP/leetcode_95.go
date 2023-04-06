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

	for i := start; i <= end; i++ {
		// 左子树集合
		leftTrees := backtrack(start, i-1)
		rightTrees := backtrack(i+1, end)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currTree := &TreeNode{i, left, right}
				allTrees = append(allTrees, currTree)
			}
		}
	}
	return allTrees
}
