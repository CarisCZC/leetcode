// https://leetcode.cn/problems/binary-tree-postorder-traversal/
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
 * @return {number[]}
 */
// 完全可借鉴144的代码
var postorderTraversal = function (root) {
  let res = [];
  if (!root) return res;
  let cur1 = root;
  let cur2 = null;
  while (cur1) {
    cur2 = cur1.left;
    if (cur2) {
      // 先找中序前驱
      while (cur2.right && cur2.right !== cur1) cur2 = cur2.right;

      // 线索化
      if (!cur2.right) {
        cur2.right = cur1;
        cur1 = cur1.left;
        continue;
      } else {
        // 此时已经回到了根
        // res.push(cur1.right);
        cur2.right = null;
        res = print(res, cur1.left);
      }
    }
    cur1 = cur1.right;
  }
  res = print(res, root);
  return res;
};
const print = function (res, root) {
  // 先翻转
  const reverseList = reverse(root);
  let cur = reverseList;
  // 再打印
  while (cur) {
    res.push(cur.val);
    cur = cur.right;
  }
  // 再逆转回来
  reverse(reverseList);
  return res;
};
const reverse = function (root) {
  let cur = root;
  let pre = null;
  while (cur) {
    const next = cur.right;
    cur.right = pre;
    pre = cur;
    cur = next;
  }
  return pre;
};
