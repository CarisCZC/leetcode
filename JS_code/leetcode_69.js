// https://leetcode.cn/problems/sqrtx/
/**
 * @param {number} x
 * @return {number}
 */
// 利用二分来找，不一定准确算出
var mySqrt = function (x) {
  if (x === 0 || x === 1) return x;
  for (let q = 0, p = x; p > q; ) {
    const mid = Math.trunc((p + q) / 2);
    const doMid = mid * mid;
    // const min = (mid - 1) * (mid - 1);
    if (doMid === x) return mid;
    if (doMid > x) {
      const tmp = (mid - 1) * (mid - 1);
      if (tmp > x) p = mid - 1;
      else return mid - 1;
    } else {
      const a = (mid + 1) * (mid + 1);
      if (a < x) q = mid + 1;
      else return a === x ? mid + 1 : mid;
    }
  }
};

// 优秀写法
// var mySqrt = function (x) {
//   let l = 0,
//     r = x,
//     ans = -1;
//   while (l <= r) {
//     const mid = Marth.trunc((l + r) / 2);
//     if (mid * mid <= x) {
//       ans = mid;
//       l = mid + 1;
//     } else {
//       r = mid - 1;
//     }
//   }
//   return ans;
// };
