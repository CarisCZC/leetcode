// https://leetcode.cn/problems/sum-root-to-leaf-numbers/
// 129. 求根节点到叶节点数字之和

package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	res := helper(root, 0, []int{})
	ans := 0
	for _, v := range res {
		ans += v
	}
	return ans
}

func helper(root *TreeNode, val int, res []int) []int {
	tmp := val*10 + root.Val
	if root.Left == nil && root.Right == nil {
		res = append(res, tmp)
		return res
	}
	if root.Left != nil {
		res = helper(root.Left, tmp, res)
	}
	if root.Right != nil {
		res = helper(root.Right, tmp, res)
	}
	return res
}
