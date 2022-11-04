// https://leetcode.cn/problems/same-tree/
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {boolean}
 */
var isSameTree = function (p, q) {
  return compare(true, p, q);
};

// 判断两个是否完全一致
const compare = function (ans, p, q) {
  // 有左右子树先下探
  // 该节点值不等，直接返回
  if (p === null && q === null) return true;
  else if (p === null || q === null) return false;
  if (p.val !== q.val) ans = false;
  if (p.left !== null && q.left !== null) ans = compare(ans, p.left, q.left);
  else if (ans && (p.left !== null || q.left !== null)) ans = false;
  if (ans && p.right !== null && q.right !== null)
    ans = compare(ans, p.right, q.right);
  else if (ans && (p.right !== null || q.right !== null)) ans = false;
  return ans;
};
// 优秀写法
// const isSameTree = function (p, q) {
//   // 都为空时
//   if (!p && !q) return true;
//   // 只有一个为空
//   else if (!p && !q) return false;
//   // 都不为空比较值
//   else if (p.val !== q.val) return false;
//   // 递归下探
//   return isSameTree(p.left, q.left) && isSameTree(p.right, q.right);
// };
