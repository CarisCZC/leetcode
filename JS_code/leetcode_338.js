//https://leetcode.cn/problems/counting-bits/
/**
 * @param {number} n
 * @return {number[]}
 */
// 1. 暴力解，一个个的算，能算出来
var countBits = function (n) {
  let res = new Array(n + 1).fill(0);
  // 动态规划
  for (let i = 1; i <= n; i++) {
    // i&(i-1),将最低为的1置为0
    // y = x&(x-1),y<x,但bit[x] = bit[y]+1
    // 即 bit[x] = bit[x&(x-1)]+1
    res[i] = res[i & (i - 1)] + 1;
  }
  return res;
};
