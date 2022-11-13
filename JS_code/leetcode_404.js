// https://leetcode.cn/problems/sum-of-left-leaves/
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
var sumOfLeftLeaves = function (root) {
  let res = 0;
  if (isLeaf(root)) return res;
  const leftLeaves = findLeftLeaves([], root);
  for (const leaf of leftLeaves) {
    res += leaf.val;
  }
  return res;
};

const findLeftLeaves = function (res, root) {
  if (!root) return res;
  // 是叶子结点
  if (isLeaf(root)) res.push(root);
  findLeftLeaves(res, root.left);
  if (!isLeaf(root.right)) findLeftLeaves(res, root.right);
  return res;
};
const isLeaf = function (root) {
  if (!root) return false;
  if (!root.left && !root.right) return true;
  return false;
};
