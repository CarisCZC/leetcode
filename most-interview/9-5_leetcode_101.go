// https://leetcode.cn/problems/symmetric-tree/
// 101. 对称二叉树
package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 对称，左子树的右节点等于右子树的左节点
func isSymmetric(root *TreeNode) bool {
	return helper(root.Left, root.Right)
}

func helper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	if !helper(left.Left, right.Right) {
		return false
	}
	if !helper(left.Right, right.Left) {
		return false
	}
	return true
}
