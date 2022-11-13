// https://leetcode.cn/problems/binary-watch/
/**
 * @param {number} turnedOn
 * @return {string[]}
 */

// 二进制手表
// 灯亮的个数，表示的是1的个数
// 小时：0-11 最多亮3个灯
// 分钟：0~59 最多亮5个灯
// 超过8个灯都无效
var readBinaryWatch = function (turnedOn) {
  const res = [];
  for (let h = 0; h < 12; h++)
    for (let m = 0; m < 60; m++) {
      if (
        h.toString(2).split("0").join("").length +
          m.toString(2).split("0").join("").length ===
        turnedOn
      )
        res.push(h + ":" + (m < 10 ? "0" : "") + m);
    }
  return res;
};
