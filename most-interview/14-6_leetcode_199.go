// https://leetcode.cn/problems/binary-tree-right-side-view/
// 199. 二叉树的右视图

package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 有右结点返回右节点，没有右结点返回左节点。
func rightSideView(root *TreeNode) []int {
	res := []int{}
	return helper(root, res, 1)
}

// res的长度代表深度
func helper(root *TreeNode, res []int, deep int) []int {
	if root == nil {
		return res
	}
	if deep > len(res) {
		res = append(res, root.Val)
	}
	res = helper(root.Right, res, deep+1)
	res = helper(root.Left, res, deep+1)
	return res
}
