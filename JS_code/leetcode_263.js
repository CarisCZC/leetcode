// https://leetcode.cn/problems/ugly-number/
/**
 * @param {number} n
 * @return {boolean}
 */
// 只包含2,3,5质因数
var isUgly = function (n) {
  if (n <= 0) return false;
  if (n === 1) return true;
  const target = [2, 3, 5];
  // 找出所有质因数
  for (let i = 1; i <= n / i; i++) {
    if (!(n % i)) {
      // 能整除，是约数
      const tmp = n / i;
      // 判断i，tmp是否是质数
      if (zhishu(i) && !target.includes(i)) return false;
      if (zhishu(tmp) && !target.includes(tmp)) return false;
    }
  }
  return true;
};

const zhishu = function (n) {
  if (n === 1) return false;
  let flag = true;
  for (let j = 2; j <= n / j; j++) {
    if (!(n % j)) {
      flag = !flag;
      break;
    }
  }
  return flag;
};

// 官方优秀案例
// var isUgly = function (n) {
//   if (n <= 0) return false;
//   const factors = [2, 3, 5];
//   for (const factor of factors) {
//     while (n % factor === 0) {
//       n /= factor;
//     }
//   }
//   return n == 1;
// };
