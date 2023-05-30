// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 104. 二叉树的最大深度

package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 最大深度
func maxDepth(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, deep int) int {
	if root == nil {
		return deep
	}
	deep++
	leftDeep := helper(root.Left, deep)
	rightDeep := helper(root.Right, deep)
	if leftDeep > rightDeep {
		deep = leftDeep
	} else {
		deep = rightDeep
	}
	return deep
}
