// https://leetcode.cn/problems/balanced-binary-tree/
// 110. 平衡二叉树
package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 和104题的思路基本一致，但要求对每个二叉树都要比较
func isBalanced(root *TreeNode) bool {
	_, isB := helper(root, 0)
	return isB
}

func helper(root *TreeNode, deep int) (int, bool) {
	if root == nil {
		return deep, true
	}
	deep++
	lfdeep, lfb := helper(root.Left, deep)
	rgdeep, rgb := helper(root.Right, deep)
	deep = max(lfdeep, rgdeep)
	if !(lfb && rgb) {
		return deep, false
	}
	// 两边都平衡，看该结点平衡
	if lfdeep-rgdeep < -1 || lfdeep-rgdeep > 1 {
		return deep, false
	}
	return deep, true

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}
