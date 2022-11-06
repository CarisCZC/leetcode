// https://leetcode.cn/problems/intersection-of-two-linked-lists/
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} headA
 * @param {ListNode} headB
 * @return {ListNode}
 */

// 1. hash的方法，能过,时间复杂度O(2n+m),空间复杂度O(n+m)
// 2. A,B都找到末尾，如果末尾结点相同，说明相交
var getIntersectionNode = function (headA, headB) {
  let curA = headA;
  let curB = headB;
  while (curA !== curB) {
    if (!curA) curA = headB;
    else curA = curA.next;
    if (!curB) curB = headA;
    else curB = curB.next;
  }
  return curA ? curA : null;
};
