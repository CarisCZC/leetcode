// https://leetcode.cn/problems/minimum-depth-of-binary-tree/
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
// 104题反过来
/**
 * @param {TreeNode} root
 * @return {number}
 */
var minDepth = function (root) {
  // 结束必须到叶子节点
  if (!root) return 0;
  if (!root.left && !root.right) return 1;
  return Math.min(
    root.left ? 1 + minDepth(root.left) : Number.MAX_VALUE,
    root.right ? 1 + minDepth(root.right) : Number.MAX_VALUE
  );
};
