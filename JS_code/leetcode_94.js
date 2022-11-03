// https://leetcode.cn/problems/binary-tree-inorder-traversal/
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
// 二叉树中序遍历
/**
 * @param {TreeNode} root
 * @return {number[]}
 */
var inorderTraversal = function (root) {
  // let res = [];
  // res = traversal(res, root);
  // if (root.left !== null) traversal(res, root.left);
  // res.push(root.val);
  // if ((root, left !== null)) traversal(res, root.right);

  return root === null ? [] : traversal([], root);
};

const traversal = function (res, root) {
  if (root.left !== null) traversal(res, root.left);
  res.push(root.val);
  if (root.right !== null) traversal(res, root.right);
  return res;
};
