// https://leetcode.cn/problems/path-sum/
// 112. 路径总和

package main

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先找到所有到叶子节点路径
// 再一一判断
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	path := helper(root, []int{}, 0)
	for i := range path {
		if targetSum == path[i] {
			return true
		}
	}
	return false
}

func helper(root *TreeNode, path []int, cur int) []int {
	if root.Left == nil && root.Right == nil {
		path = append(path, cur+root.Val)
		return path
	}
	if root.Left != nil {
		path = helper(root.Left, path, root.Val+cur)
	}
	if root.Right != nil {
		path = helper(root.Right, path, root.Val+cur)
	}
	return path
}
