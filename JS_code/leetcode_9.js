// https://leetcode.cn/problems/palindrome-number/
/**
 * @param {number} x
 * @return {boolean}
 */
// 双坐标
var isPalindrome = function (x) {
  const xStr = "" + x;
  for (let i = 0, j = xStr.length - 1; i < xStr.length / 2; i++, j--) {
    if (xStr[i] !== xStr[j]) return false;
  }
  return true;
};

// console.log(isPalindrome(1221));
