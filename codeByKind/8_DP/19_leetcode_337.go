package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 从root进入，左子树右子树互不影响。
// 可以分，偷root的情况，和不偷root的情况
// 即dp[0]表示不偷当前结点的情况，dp[1]表示偷该节点所得到的最大金钱
func rob(root *TreeNode) int {
	res := robTree(root)
	return max(res[0], res[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robTree(cur *TreeNode) []int {
	if cur == nil {
		return []int{0, 0}
	}

	left := robTree(cur.Left)
	right := robTree(cur.Right)
	notRobCur := max(left[0], left[1]) + max(right[0], right[1])
	robCur := cur.Val + left[0] + right[0]
	return []int{notRobCur, robCur}
}
