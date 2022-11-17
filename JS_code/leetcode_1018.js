// https://leetcode.cn/problems/binary-prefix-divisible-by-5/
/**
 * @param {number[]} nums
 * @return {boolean[]}
 */
var prefixesDivBy5 = function (nums) {
  const res = [];
  let cur = 0;
  for (const num of nums) {
    cur = ((cur << 1) + num) % 5;
    res.push(cur === 0);
  }
  return res;
};
