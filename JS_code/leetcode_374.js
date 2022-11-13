//https://leetcode.cn/problems/guess-number-higher-or-lower/

/**
 * Forward declaration of guess API.
 * @param {number} num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * var guess = function(num) {}
 */

/**
 * @param {number} n
 * @return {number}
 */

// 就是二分查找
var guessNumber = function (n) {
  let left = 1,
    right = n;
  mid = Math.trunc((left + right) / 2);
  while (left < right) {
    if (!guess(mid)) return mid;
    if (guess(mid) > 0) left = mid + 1;
    else right = mid - 1;
    mid = Math.trunc((left + right) / 2);
  }
  return mid;
};
