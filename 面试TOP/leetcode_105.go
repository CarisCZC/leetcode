package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 先找出root
	var root TreeNode
	if len(preorder) == 0 {
		return nil
	}
	root.Val = preorder[0]
	if len(preorder) == 1 {
		return &root
	}
	index := 0
	for i, v := range inorder {
		if v == root.Val {
			index = i
			break
		}
	}
	//构造左子树，构造右子树
	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return &root
}
