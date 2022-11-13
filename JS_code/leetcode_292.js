// https://leetcode.cn/problems/nim-game/
/**
 * @param {number} n
 * @return {boolean}
 */
// Nim 游戏
// 一人一次，消耗2-6个石头
// 必赢情况：最后到我时，只剩下1-3个石头
// 只要是4的倍数，无论自己怎么拿，对方一定能控制，最后剩下4个自己拿不下
var canWinNim = function (n) {
  return n % 4 === 0;
};
