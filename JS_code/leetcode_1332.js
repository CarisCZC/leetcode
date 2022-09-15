// 1332. 删除回文子序列
// https://leetcode.cn/problems/remove-palindromic-subsequences/
/**
 * @param {string} s
 * @return {number}
 */

// 如果是回文串，返回1
// 如果不是回文串，返回2
// 没意思这题
var removePalindromeSub = function (s) {
  let i = 0,
    j = s.length - 1;
  while (i < j) {
    if (s[i] !== s[j]) return 2;
    i++;
    j--;
  }
  return 1;
};
