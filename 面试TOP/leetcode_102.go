// https://leetcode.cn/problems/binary-tree-level-order-traversal/?favorite=2ckc81c

package main

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	thisLevel := []*TreeNode{}
	res := [][]int{}
	if root == nil {
		return res
	}
	if root != nil {
		res = append(res, []int{root.Val})
	}
	levelcnt := 0
	if root.Left != nil {
		thisLevel = append(thisLevel, root.Left)
		levelcnt++
	}
	if root.Right != nil {
		thisLevel = append(thisLevel, root.Right)
		levelcnt++
	}
	nextLevelCnt := 0
	cur := []int{}
	for levelcnt != 0 {

		curNode := thisLevel[0]
		if curNode.Left != nil {
			thisLevel = append(thisLevel, curNode.Left)
			nextLevelCnt++
		}
		if curNode.Right != nil {
			thisLevel = append(thisLevel, curNode.Right)
			nextLevelCnt++
		}
		cur = append(cur, thisLevel[0].Val)
		thisLevel = thisLevel[1:]
		levelcnt--
		if levelcnt == 0 {
			levelcnt = nextLevelCnt
			nextLevelCnt = 0
			res = append(res, cur)
			cur = []int{}
		}

	}
	return res
}
