// leetcode 257， 二叉树的所有路径

// 给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

package main

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var paths = []string{}

func binaryTreePaths(root *TreeNode) []string {
	findPath(root, "")
	return paths
}

func findPath(root *TreeNode, path string) {
	if root != nil {
		pathSB := path
		pathSB += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			paths = append(paths, pathSB)
		} else {
			pathSB += "->"
			findPath(root.Left, pathSB)
			findPath(root.Right, pathSB)
		}
	}
}
