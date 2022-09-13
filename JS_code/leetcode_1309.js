// 1309. 解码字母到整数映射
// https://leetcode.cn/problems/decrypt-string-from-alphabet-to-integer-mapping/
/**
 * @param {string} s
 * @return {string}
 */
var freqAlphabets = function (s) {
  const letters = "abcdefghijklmnopqrstuvwxyz";
  let res = "";
  for (let i = s.length - 1; i > -1; ) {
    if (s[i] === "#") {
      // 往前走两位

      res = letters[s.slice(i - 2, i) - 1] + res;
      i -= 3;
      continue;
    }
    res = letters[s.slice(i, i + 1) - 1] + res;
    i--;
  }
  return res;
};
