/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}
	res := 0
	// 遇到树，一定要递归
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int { // root是cur的父节点,isWatch 表示，该节点是否被监控
		// 由于我们是从父节点到子节点，因此，该节点如果被监控，那么一定是被父节点监控。
		// 自底向上遍历，最合适
		if root.Left == nil && root.Right == nil { //到子节点了
			return 1 // 请求父节点监视自己
		} else {
			// 1 本节点没有放置监视器，且需要父节点放置
			// 0 本节点没有放置，但也被监视
			// -1 本节点点放置了监视器
			leftWatch, rightWatch := -2, -2
			if root.Left != nil {
				leftWatch = dfs(root.Left)
			}
			if root.Right != nil {
				rightWatch = dfs(root.Right)
			}
			if leftWatch > 0 || rightWatch > 0 { //
				res++
				return -1
			}
			if leftWatch == -1 || rightWatch == -1 {
				// leftWatch == -1 且 rightWatch == -1
				return 0
			}
			if leftWatch == 0 || rightWatch == 0 {
				// 返回给其父节点的，其子节点都监视了，但其自己还没有被监视
				return 1
			}
			return -2
		}
	}
	needWatch := dfs(root)
	if needWatch > 0 {
		res++
	}
	return res
}
