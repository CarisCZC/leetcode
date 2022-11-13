//https://leetcode.cn/problems/palindrome-linked-list/
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {boolean}
 */
/**
 * 回文链表
 * 1. 用一个数组来承接
 * 2. 转变成回文字符串判断
 * 3. 快慢指针（答题以此为主）
 */
var isPalindrome = function (head) {
  if (!head) return true;
  // 分半
  let firstHalfEnd = endOfFirstHalf(head);
  // 翻转后半
  let secondHalfStart = reverseList(firstHalfEnd.next);
  // 回文判断
  let p1 = head;
  let p2 = secondHalfStart;
  let res = true;
  while (res && p2) {
    if (p1.val !== p2.val) res = false;
    p1 = p1.next;
    p2 = p2.next;
  }
  // 恢复
  firstHalfEnd.next = reverseList(secondHalfStart);
  return res;
};

// 翻转后半链表
const reverseList = function (head) {
  let prev = null;
  let curr = head;
  while (curr) {
    let tmp = curr.next;
    curr.next = prev;
    prev = curr;
    curr = tmp;
  }
  return prev;
};
// 快慢指针找后半链表
const endOfFirstHalf = function (head) {
  let fast = head;
  let slow = head;
  while (fast.next && fast.next.next) {
    fast = fast.next.next;
    slow = slow.next;
  }
  return slow;
};
