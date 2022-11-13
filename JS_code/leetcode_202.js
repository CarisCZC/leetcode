// https://leetcode.cn/problems/happy-number/
/**
 * @param {number} n
 * @return {boolean}
 */
var isHappy = function (n) {
  let slow = n;
  let fast = getNext(n);
  while (fast !== 1 && slow != fast) {
    slow = getNext(slow);
    fast = getNext(getNext(fast));
  }
  return fast === 1;
};

const getNext = function (n) {
  let totalSum = 0;
  while (n > 0) {
    let d = n % 10;
    n = Math.trunc(n / 10);
    totalSum += d ** 2;
  }
  return totalSum;
};
