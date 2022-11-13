//https://leetcode.cn/problems/remove-linked-list-elements/
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} val
 * @return {ListNode}
 */
var removeElements = function (head, val) {
  const newHead = new ListNode(0, head);
  let q = newHead;
  while (q.next) {
    if (q.next.val === val) {
      q.next = q.next.next;
    } else {
      q = q.next;
    }
  }
  return newHead.next;
};
