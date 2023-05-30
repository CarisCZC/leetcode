package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	// 有个队列，放入遍历的结点
	nodeList := []*TreeNode{}
	res := [][]int{}
	if root == nil {
		return res
	}
	nodeList = append(nodeList, root)
	for len(nodeList) != 0 {
		cur := []int{}
		nodeList, cur = levelHpler(nodeList)
		res = append(res, cur)
	}
	return res
}

func levelHpler(nodeList []*TreeNode) ([]*TreeNode, []int) {
	levelLen := len(nodeList)
	res := []int{}
	for i := 0; i < levelLen; i++ {
		cur := nodeList[i]
		res = append(res, cur.Val)
		if cur.Left != nil {
			nodeList = append(nodeList, cur.Left)
		}
		if cur.Right != nil {
			nodeList = append(nodeList, cur.Right)
		}
	}
	nodeList = nodeList[levelLen:]
	return nodeList, res
}
