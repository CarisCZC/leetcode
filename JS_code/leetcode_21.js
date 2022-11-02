//https://leetcode.cn/problems/merge-two-sorted-lists/
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} list1
 * @param {ListNode} list2
 * @return {ListNode}
 */
var mergeTwoLists = function (list1, list2) {
  let p = list1,
    q = list2;
  const res = { val: 0, next: null };
  let r = res;
  while (p !== null && q !== null) {
    if (p.val <= q.val) {
      r.next = p;
      p = p.next;
    } else {
      r.next = q;
      q = q.next;
    }
    r = r.next;
  }
  if (p !== null) {
    r.next = p;
  }
  if (q !== null) {
    r.next = q;
  }
  //r.next = p == null?q:p;
  return res.next;
};
