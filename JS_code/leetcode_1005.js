// https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations/
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var largestSumAfterKNegations = function (nums, k) {
  // 记录一个最小非负数
  let min = Number.MAX_SAFE_INTEGER,
    cnt = 0;
  //  存放负数
  let others = [];
  for (const num of nums) {
    if (num >= 0) {
      min = Math.min(num, min);
      cnt += num;
    } else {
      others.push(num);
    }
  }
  others.sort((a, b) => a - b);
  while (others.length > 0) {
    if (k > 0) {
      let tmp = -others.shift();
      min = Math.min(tmp, min);
      cnt += tmp;
      k--;
    } else {
      cnt += others.shift();
    }
  }
  // k>0,所有负数都变正了，且剩下的k为奇数
  if (k > 0 && k % 2 === 1) cnt -= 2 * min;
  return cnt;
};

console.log(largestSumAfterKNegations([4, 2, 3]));
