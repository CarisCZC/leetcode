//  https://leetcode.cn/problems/number-of-1-bits/
/**
 * @param {number} n - a positive integer
 * @return {number}
 */
var hammingWeight = function (n) {
  let rev = 0;
  for (let i = 0; i < 32; i++) {
    if (n & 1) {
      // 如果最后一位是1
      rev++;
    }
    n >>>= 1;
  }
  return rev;
};
