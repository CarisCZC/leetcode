package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	NodeList := make([]*Node, 0)
	randomIndex := make([]int, 0)
	newNodeList := make([]*Node, 0)
	newHead := new(Node)
	q := newHead
	p := head
	headIndex := 0
	for p != nil {
		// 为了能快速找到节点，将节点转换为数组形式
		NodeList = append(NodeList, p)
		randomIndex = append(randomIndex, -1)
		p = p.Next
	}
	p = head
	for p != nil {
		// 对每个节点，都遍历一下，来判断其Random节点
		for i, random := range NodeList {
			if random == p.Random {
				// 当前节点的Random被找到
				randomIndex[headIndex] = i
			}
		}
		q.Val = p.Val
		if p.Next != nil {
			nextNode := new(Node)
			q.Next = nextNode
		}
		newNodeList = append(newNodeList, q)
		q = q.Next
		p = p.Next
		headIndex++
	}
	// 初始化random
	q = newHead
	headIndex = 0
	for q != nil {
		index := randomIndex[headIndex]
		if index == -1 {
			q.Random = nil
		} else {
			q.Random = newNodeList[index]
		}
		q = q.Next
		headIndex++
	}
	return newHead
}
