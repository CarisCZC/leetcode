// leetcode 108,将有序数组转换成二叉搜索树
package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//二分查找呗，分治
func sortedArrayToBST(nums []int) (binaryTree *TreeNode) {
	binaryTree = new(TreeNode)
	left := 0
	right := len(nums) - 1
	if left == right { //此时只有一个元素
		binaryTree.Val = nums[right]
		return
	}
	mid := (left + right) >> 1
	binaryTree.Val = nums[mid]
	if mid != left { // mid ==left说明没有左子树
		binaryTree.Left = sortedArrayToBST(nums[left:mid])
	}
	// 因为向下取整，所以一定会有右子树
	// 如果向上取整，那么可能没有右子树，但有左子树
	binaryTree.Right = sortedArrayToBST(nums[mid+1 : right+1])
	return
}

// func main() {
// 	nums := []int{-10, -3, 0, 5, 9}
// 	sortedArrayToBST(nums)
// }
