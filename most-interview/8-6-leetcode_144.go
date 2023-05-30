// https://leetcode.cn/problems/binary-tree-preorder-traversal/
// 144. 二叉树的前序遍历

package main

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	res = helper(root, res)
	return res
}

func helper(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	if root.Left != nil {
		res = helper(root.Left, res)
	}
	if root.Right != nil {
		res = helper(root.Right, res)
	}
	return res
}
