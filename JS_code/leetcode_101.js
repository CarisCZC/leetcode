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
 * @return {boolean}
 */
var isSymmetric = function (root) {
  // 如果root为空
  if (!root) return true;
  return compare(root.left, root.right);
};
const compare = function (a, b) {
  if (!a & !b) return true;
  else if (!a || !b) return false;
  else if (a.val !== b.val) return false;
  return compare(a.left, b.right) && compare(a.right, b.left);
};
