// https://leetcode.cn/problems/invert-binary-tree/
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {TreeNode}
 */
var invertTree = function (root) {
  reverse(root);
  return root;
};
const reverse = function (root) {
  if (!root) return;
  reverse(root.left);
  reverse(root.right);
  let tmp = root.left;
  root.left = root.right;
  root.right = tmp;
};
