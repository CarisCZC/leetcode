// leetcode 905 按奇偶顺序排列数组
/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortArrayByParity = function (nums) {
  // 一个从前找奇数，一个从后找偶数，奇偶交换
  let i = 0,
    j = nums.length - 1;
  while (i < j) {
    while (i < j && nums[i] % 2 === 0) i++;

    while (i < j && nums[j] % 2 === 1) j--;
    const tmp = nums[i];
    nums[i] = nums[j];
    nums[j] = tmp;
    i++;
    j--;
  }
  return nums;
};
