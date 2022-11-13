// https://leetcode.cn/problems/first-bad-version/
/**
 * Definition for isBadVersion()
 *
 * @param {integer} version number
 * @return {boolean} whether the version is bad
 * isBadVersion = function(version) {
 *     ...
 * };
 */

/**
 * @param {function} isBadVersion()
 * @return {function}
 */
var solution = function (isBadVersion) {
  /**
   * @param {integer} n Total versions
   * @return {integer} The first bad version
   */
  // 二分查找
  // bad <= n,且n一定是错误版本
  return function (n) {
    let left = 1,
      right = n;
    let mid = Math.trunc((left + right) / 2);
    while (left < right) {
      // isBadVersion(mid)=true 说明在first在右侧。否则在左侧
      let com = isBadVersion(mid);
      if (com === isBadVersion(mid - 1)) {
        if (com) right = mid - 1;
        else left = mid + 1;
      } else {
        return mid;
      }
      mid = Math.trunc((left + right) / 2);
    }
    return mid;
  };
};
