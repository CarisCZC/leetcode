// https://leetcode.cn/problems/n-repeated-element-in-size-2n-array/
// 961. 在长度 2N 的数组中找出重复 N 次的元素
/**
 * @param {number[]} nums
 * @return {number}
 */

var repeatedNTimes = function (nums) {
  const n = nums.length;
  while (true) {
    const x = Math.floor(Math.random() * n),
      y = Math.floor(Math.random() * n);
    if (x !== y && nums[x] === nums[y]) {
      return nums[x];
    }
  }
};
