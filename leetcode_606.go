// leetcode 606, 根据二叉树创建字符串

// 给你二叉树的根节点 root ，请你采用前序遍历的方式，
// 将二叉树转化为一个由括号和整数组成的字符串，返回构造出的字符串。

// // 空节点使用一对空括号对 "()" 表示，
// 转化后需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。

package main

import (
	"strconv"
	"strings"
)

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

//前序遍历

func tree2str(root *TreeNode) string {
	res := strings.Builder{}
	res.WriteString(strconv.Itoa(root.Val))
	if root.Left == nil && root.Right != nil {
		res.WriteString("()")
	}
	if root.Left != nil {
		res.WriteString("(" + tree2str(root.Left) + ")")
	}
	if root.Right != nil {
		res.WriteString("(" + tree2str(root.Right) + ")")
	}
	return res.String()

}
