//https://leetcode.cn/problems/reverse-bits/
/**
 * @param {number} n - a positive integer
 * @return {number} - a positive integer
 */
var reverseBits = function (n) {
  let rev = 0;
  for (let i = 0; i < 32; ++i) {
    // 结果先左移一位
    rev <<= 1;
    // 或上n的最后一位
    rev |= n & 1;
    // n 右移
    n >>>= 1;
  }
  return rev >>> 0;
};

// 分治法
var reverseBits = function (n) {
  const M1 = 0x55555555; // 0101 0101 0101 0101 0101 0101 0101 0101
  const M2 = 0x33333333; // 0011 0011 0011 0011 0011 0011 0011 0011
  const M3 = 0x0f0f0f0f; // 0000 1111 0000 1111 0000 1111 0000 1111
  const M4 = 0x0000ffff; // 0000 0000 0000 0000 1111 1111 1111 1111
  // M1表示保留所有奇数位，n>>>1 将偶数位变成了奇数位。
  // (n >>> 1) & M1 等于在偶数位置上保留原先奇数位的数字
  // (n & M1) << 1，在保留奇数位后，左移一位，将奇数位变成了偶数位
  n = ((n >>> 1) & M1) | ((n & M1) << 1);
  // 同理
  n = ((n >>> 2) & M2) | ((n & M2) << 2);
  n = ((n >>> 4) & M3) | ((n & M3) << 4);
  n = ((n >>> 8) & M4) | ((n & M4) << 8);
  // >>> 0 去除符号
  return (n >>> 16) | ((n << 16) >>> 0);
};
