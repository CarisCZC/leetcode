//  https://leetcode.cn/problems/power-of-four/
/**
 * @param {number} n
 * @return {boolean}
 */
var isPowerOfFour = function (n) {
  const mark = 0xaaaaaaaa;
  return n > 0 && !(n & (n - 1)) && !(n & mark);
};
