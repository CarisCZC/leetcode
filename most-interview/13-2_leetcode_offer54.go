// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
// 剑指 Offer 54. 二叉搜索树的第k大节点
package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历的顺序来找
func kthLargest(root *TreeNode, k int) int {
	cur := 0
	res := 0
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Right)
		cur++
		if cur == k {
			res = root.Val
			return
		}
		helper(root.Left)
	}
	helper(root)
	return res
}
