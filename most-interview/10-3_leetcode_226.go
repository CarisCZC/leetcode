// https://leetcode.cn/problems/invert-binary-tree/
// 226. 翻转二叉树

package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 左子树的与右子树交换
func invertTree(root *TreeNode) *TreeNode {
	return helper(root)
}

func helper(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := helper(root.Left)
	right := helper(root.Right)
	root.Left = right
	root.Right = left
	return root
}
