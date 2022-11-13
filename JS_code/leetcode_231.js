// https://leetcode.cn/problems/power-of-two/
/**
 * @param {number} n
 * @return {boolean}
 */
// 如果是2的幂，那么其二进制一定是10....0的形式
var isPowerOfTwo = function (n) {
  return n > 0 && (n & (n - 1)) === 0;
};
