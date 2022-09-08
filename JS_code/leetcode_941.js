// leetcode 941. 有效的山脉数组
// https://leetcode.cn/problems/valid-mountain-array/
/**
 * @param {number[]} arr
 * @return {boolean}
 */
// 从前往后递增，从后往前递增
var validMountainArray = function (arr) {
  if (arr.length < 3) return false;
  let i = 0,
    j = arr.length - 1;
  while (i < j) {
    while (i < j && arr[i] < arr[i + 1]) i++;
    while (i < j && arr[j] < arr[j - 1]) j--;
    return i !== 0 && j !== arr.length - 1 && i === j ? true : false;
  }
};
