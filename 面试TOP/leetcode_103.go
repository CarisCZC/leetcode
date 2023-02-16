package main

//  * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	// 是从左到右还是从右到左

	res := [][]int{}
	if root == nil {
		return res
	}
	if root != nil {
		res = append(res, []int{root.Val})
	}
	levelcnt := 0
	thisLevel := []*TreeNode{}
	if root.Right != nil {
		thisLevel = append(thisLevel, root.Right)
		levelcnt++
	}
	if root.Left != nil {
		thisLevel = append(thisLevel, root.Left)
		levelcnt++
	}
	LToR := true
	nextLevelCnt := 0
	cur := []int{}
	tmp := []*TreeNode{}
	for levelcnt != 0 {
		curNode := thisLevel[0]
		cur = append(cur, curNode.Val)
		thisLevel = thisLevel[1:]
		tmp = append(tmp, curNode)
		levelcnt--
		if levelcnt == 0 {
			if LToR {
				for i := len(tmp) - 1; i > -1; i-- {
					if tmp[i].Left != nil {
						thisLevel = append(thisLevel, tmp[i].Left)
						nextLevelCnt++
					}
					if tmp[i].Right != nil {
						thisLevel = append(thisLevel, tmp[i].Right)
						nextLevelCnt++
					}
				}
				LToR = !LToR
			} else {
				for i := len(tmp) - 1; i > -1; i-- {
					if tmp[i].Right != nil {
						thisLevel = append(thisLevel, tmp[i].Right)
						nextLevelCnt++
					}
					if tmp[i].Left != nil {
						thisLevel = append(thisLevel, tmp[i].Left)
						nextLevelCnt++
					}
				}
				LToR = !LToR
			}
			tmp = []*TreeNode{}
			levelcnt = nextLevelCnt
			nextLevelCnt = 0
			res = append(res, cur)
			cur = []int{}
		}
	}
	return res
}
