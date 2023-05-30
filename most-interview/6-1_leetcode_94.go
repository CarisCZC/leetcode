// https://leetcode.cn/problems/binary-tree-inorder-traversal/
// 94. 二叉树的中序遍历

package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归式遍历
func inorderTraversal(root *TreeNode) []int {
	return helper(root, []int{})
}

func helper(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = helper(root.Left, res)
	res = append(res, root.Val)
	res = helper(root.Right, res)
	return res
}
