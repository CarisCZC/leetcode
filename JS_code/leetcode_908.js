// leetcode 908 最小差值
// https://leetcode.cn/problems/smallest-range-i/
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
// 先找到最小值和最大值，然后对最小值和最大值进行调整
var smallestRangeI = function (nums, k) {
  let min = Number.MAX_SAFE_INTEGER,
    max = -1;
  nums.forEach((num) => {
    if (num < min) min = num;
    if (num > max) max = num;
  });
  //   console.log(min, max);
  const chazhi = max - min;
  if (k * 2 >= chazhi) return 0;
  return chazhi - k * 2;
};
console.log(smallestRangeI([1, 2, 4, 7], 2));
