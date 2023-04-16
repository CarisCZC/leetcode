// https://leetcode.cn/problems/binary-tree-maximum-path-sum/?favorite=2ckc81c

package main

import "math"

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxRoad := math.MinInt32
	var findRoad func(*TreeNode) int
	findRoad = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftRoad := max(findRoad(root.Left), 0)
		rightRoad := max(findRoad(root.Right), 0)
		priceNewPath := root.Val + leftRoad + rightRoad
		maxRoad = max(maxRoad, priceNewPath)
		return root.Val + max(leftRoad, rightRoad)
	}
	findRoad(root)
	return maxRoad
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
