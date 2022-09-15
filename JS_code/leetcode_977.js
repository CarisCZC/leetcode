// 977. 有序数组的平方
// https://leetcode.cn/problems/squares-of-a-sorted-array/
/**
 * @param {number[]} nums
 * @return {number[]}
 */
// 从正负中间开始分断，左右两部分
var sortedSquares = function (nums) {
  let flag = -1;
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] < 0 && (nums[i + 1] >= 0 || i === nums.length - 1)) {
      flag = i;
      break;
    }
  }
  const res = [];
  let i = flag,
    j = flag + 1;
  while (i > -1 && j < nums.length) {
    const a = nums[i] ** 2,
      b = nums[j] ** 2;
    if (a <= b) {
      res.push(a);
      i--;
    } else {
      res.push(b);
      j++;
    }
  }
  while (i > -1) {
    res.push(nums[i] ** 2);
    i--;
  }
  while (j < nums.length) {
    res.push(nums[j] ** 2);
    j++;
  }
  return res;
};

console.log(sortedSquares([-5, -3, -2, -1]));
