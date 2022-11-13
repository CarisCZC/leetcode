// https://leetcode.cn/problems/reverse-linked-list/
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var reverseList = function (head) {
  if (!head) return head;
  let q = new ListNode(head.val);
  q.next = null;
  let p = head.next;
  while (p) {
    let tmp = p.next;
    p.next = q;
    q = p;
    p = tmp;
  }
  return q;
};
