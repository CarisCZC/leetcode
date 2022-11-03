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
var deleteDuplicates = function (head) {
  let prv = head;
  const first = head;
  while (prv !== null && prv.next != null) {
    const tmpV = prv.val;
    if (prv.next?.val === tmpV) prv.next = prv.next.next;
    else prv = prv.next;
  }
  return first;
};
