//  https://leetcode.cn/problems/valid-perfect-square/
/**
 * @param {number} num
 * @return {boolean}
 */

var isPerfectSquare = function (num) {
  if (num === 1) return true;
  for (let i = 2; i <= num / 2; i++) {
    if (i ** 2 === num) return true;
  }
  return false;
};
