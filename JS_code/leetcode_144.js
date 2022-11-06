// https://leetcode.cn/problems/binary-tree-preorder-traversal/
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
// morris 遍历 即线索化
// 前序
var preorderTraversal = function (root) {
  const res = [];
  if (!root) return res;
  let cur1 = root;
  let cur2 = null;
  while (cur1) {
    cur2 = cur1.left;
    if (cur2) {
      // 找到cur1的中序前驱
      while (cur2.right && cur2.right !== cur1) cur2 = cur2.right;
      // cur2不存在右子树，将cur1绑定到cur2的右子树
      if (!cur2.right) {
        cur2.right = cur1;
        // 将cur1输出
        res.push(cur1.val);
        // cur1 前进
        cur1 = cur1.left;
        continue;
      } else {
        //当到达叶子时，cur2.right存在&&cur2.right=0，为了防止重复回来，cur2，right把链断开
        cur2.right = null;
      }
    } else {
      res.push(cur1.val);
    }
    cur1 = cur1.right;
  }
  return res;
};
