// 1189. “气球” 的最大数量
// https://leetcode.cn/problems/maximum-number-of-balloons/

/**
 * @param {string} text
 * @return {number}
 */
var maxNumberOfBalloons = function (text) {
  // b a l o n
  const cnt = new Array(5).fill(0);
  for (const ch of text) {
    if (ch === "b") {
      cnt[0]++;
    } else if (ch === "a") {
      cnt[1]++;
    } else if (ch === "l") {
      cnt[2] += 0.5;
    } else if (ch === "o") {
      cnt[3] += 0.5;
    } else if (ch === "n") {
      cnt[4]++;
    }
  }
  console.log(cnt);
  return Math.floor(Math.min(...cnt));
};

console.log(maxNumberOfBalloons("balon"));
