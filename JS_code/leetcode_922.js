// 922. 按奇偶排序数组 II
// https://leetcode.cn/problems/sort-array-by-parity-ii/

// 一个从偶数开始找奇数，一个从奇数开始找偶数，交换
/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortArrayByParityII = function (nums) {
  /* 
  除了酷毫无用处
    const odd = [],
      even = [];
    return nums
      .map((value) => (value % 2 === 0 ? even.push(value) : odd.push(value)))
      .map((_, i) => (i % 2 === 0 ? even.shift() : odd.shift()));
  */
  // 偶数下标和奇数下标
  let even = 0,
    odd = 1;
  while (odd < nums.length && even < nums.length) {
    if (nums[odd] % 2 === 0 && nums[even] % 2 !== 0)
      [nums[odd], nums[even]] = [nums[even], nums[odd]];
    if (nums[odd] % 2 !== 0) odd += 2;
    if (nums[even] % 2 === 0) even += 2;
  }
  return nums;
};

console.log(sortArrayByParityII([1, 2, 3, 4, 5, 6, 7, 8]));
