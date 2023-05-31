// https://leetcode.cn/problems/diameter-of-binary-tree/
// 543. 二叉树的直径
package main

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 左右子树的深度相加
func diameterOfBinaryTree(root *TreeNode) int {
	l, _ := helper(root, 0)
	return l
}

// 两个参数，第一个参数是当前节点的最大直径。第二个参数是当前节点的左右最大深度
func helper(root *TreeNode, deep int) (int, int) {
	if root == nil {
		return 0, deep - 1
	}
	maxl, deepl := helper(root.Left, deep+1)
	maxr, deepr := helper(root.Right, deep+1)

	l := 0
	if deepl+deepr-deep*2 > max(maxr, maxl) {
		l = deepl + deepr - deep*2
	} else {
		l = max(maxr, maxl)
	}
	return l, max(deepl, deepr)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
