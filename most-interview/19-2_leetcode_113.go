// https://leetcode.cn/problems/path-sum-ii/
// 113. 路径总和 II

package main

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	res = helper(root, 0, targetSum, res, []int{})
	return res
}

func helper(root *TreeNode, curSum, targetSum int, res [][]int, path []int) [][]int {
	path = append(path, root.Val)
	curSum = curSum + root.Val
	if root.Left == nil && root.Right == nil {
		if curSum == targetSum {
			tmp := []int{}
			tmp = append(tmp, path...)
			res = append(res, tmp)
		}
		return res
	}
	if root.Left != nil {
		res = helper(root.Left, curSum, targetSum, res, path)
	}
	if root.Right != nil {
		res = helper(root.Right, curSum, targetSum, res, path)
	}
	return res
}
