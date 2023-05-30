// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
// 236. 二叉树的最近公共祖先

package main

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 将祖先结点放入队列。找最后出现的共同祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, pathp := find(root, p, []*TreeNode{})
	_, pathq := find(root, q, []*TreeNode{})
	var res *TreeNode
	for len(pathp) != 0 && len(pathq) != 0 {
		if pathp[0] == pathq[0] {
			res = pathp[0]
		}
		pathp = pathp[1:]
		pathq = pathq[1:]
	}
	return res
}

// 从根节点找到目标结点，所经历过的结点
func find(root, target *TreeNode, path []*TreeNode) (bool, []*TreeNode) {
	if root == nil {
		return false, path
	}
	path = append(path, root)
	if root == target {
		return true, path
	}
	isfind, lPath := find(root.Left, target, path)
	if isfind {
		return true, lPath
	}
	isfind, rpath := find(root.Right, target, path)
	if isfind {
		return true, rpath
	}
	return false, path
}
