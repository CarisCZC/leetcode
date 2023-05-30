// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/
// 103. 二叉树的锯齿形层序遍历

package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 也是层次遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	isLeft := true // true表示从左向右遍历。false 表示从右向左
	nodeQueue := []*TreeNode{}
	thisLevel := []int{}
	nodeQueue = append(nodeQueue, root)
	for len(nodeQueue) != 0 {
		nodeQueue, thisLevel = helper(nodeQueue, isLeft)
		isLeft = !isLeft
		res = append(res, thisLevel)
	}
	return res
}

func helper(nodeQueue []*TreeNode, isLeft bool) ([]*TreeNode, []int) {
	cnt := len(nodeQueue)
	res := []int{}
	for i := 0; i < cnt; i++ {
		cur := nodeQueue[i]
		res = append(res, cur.Val)
		if cur.Left != nil {
			nodeQueue = append(nodeQueue, cur.Left)
		}
		if cur.Right != nil {
			nodeQueue = append(nodeQueue, cur.Right)
		}
	}
	nodeQueue = nodeQueue[cnt:]
	if !isLeft { // 逆置res
		for i, j := 0, len(res)-1; i < j; {
			res[i], res[j] = res[j], res[i]
			i++
			j--
		}
	}
	return nodeQueue, res
}
