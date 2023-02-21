// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/?favorite=2ckc81c

package main

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 得层次遍历
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	thisLevel := []*Node{}
	levelCnt := 0
	// root结点不需要加*next，从第二层开始
	if root.Left != nil {
		thisLevel = append(thisLevel, root.Left)
		levelCnt++
	}
	if root.Right != nil {
		thisLevel = append(thisLevel, root.Right)
		levelCnt++
	}
	// pre, cur := &Node{}, &Node{}
	for levelCnt != 0 {
		var pre *Node
		for i := 0; i < levelCnt; i++ {
			if i == 0 {
				pre = thisLevel[i]
			} else {
				pre.Next = thisLevel[i]
				pre = pre.Next
			}
			if pre.Left != nil {
				thisLevel = append(thisLevel, pre.Left)
			}
			if pre.Right != nil {
				thisLevel = append(thisLevel, pre.Right)
			}
		}
		thisLevel = thisLevel[levelCnt:]
		levelCnt = len(thisLevel)
	}
	return root
}
