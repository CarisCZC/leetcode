// https://leetcode.cn/problems/add-digits/
// 各数相加
/**
 * @param {number} num
 * @return {number}
 */
// 与快乐数有些类似，快乐数的难点在于判断无限循环
var addDigits = function (num) {
  if (num < 10) return num;
  let res = num;
  while (!(res < 10)) {
    let tmp = res;
    res = 0;
    while (tmp) {
      res += tmp % 10;
      tmp = Math.trunc(tmp / 10);
    }
  }
  return res;
};
