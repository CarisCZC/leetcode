// https://leetcode.cn/problems/validate-binary-search-tree/?favorite=2ckc81c
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先中序遍历，然后判断遍历数组是否是升序
func isValidBST(root *TreeNode) bool {
	res := inOrder(root, []int{})
	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode, res []int) []int {
	if root.Left == nil && root.Right == nil {
		res = append(res, root.Val)
		return res
	}
	if root.Left != nil {
		res = inOrder(root.Left, res)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = inOrder(root.Right, res)
	}
	return res
}
