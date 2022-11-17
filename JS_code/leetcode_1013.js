// https://leetcode.cn/problems/partition-array-into-three-parts-with-equal-sum/
/**
 * @param {number[]} arr
 * @return {boolean}
 */
var canThreePartsEqualSum = function (arr) {
  // 先计算和，判断和是否是3的倍数
  let cnt = 0;
  for (const num of arr) {
    cnt += num;
  }
  if (cnt % 3) return false;
  else {
    cnt /= 3;
  }
  let index = [];
  let tmp = 0;
  for (let i = 0; i < arr.length; i++) {
    tmp += arr[i];
    if (tmp === cnt) {
      index.push(i);
      tmp = 0;
    }
    if (index.length === 2) {
      break;
    }
  }
  if (index.length < 2 || index[1] === arr.length - 1) return false;
  return true;
};
