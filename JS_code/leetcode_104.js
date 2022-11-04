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
 * @return {number}
 */
var maxDepth = function (root) {
  return comDepth(0, root);
};

const comDepth = function (num, root) {
  // 到达空结点
  if (!root) return num;
  return max(comDepth(num + 1, root.left), comDepth(num + 1, root.right));
};
const max = function (a, b) {
  if (a > b) return a;
  return b;
};
