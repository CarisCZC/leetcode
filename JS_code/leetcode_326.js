// https://leetcode.cn/problems/power-of-three/
/**
 * @param {number} n
 * @return {boolean}
 */

// 3的幂
// 1. 循环判断
// 2. 枚举（取巧就找最大的3的幂）
var isPowerOfThree = function (n) {
  // 找最大的3的幂，然后判断
  return n > 0 && 1162261467 % n === 0;
};
