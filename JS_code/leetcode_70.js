// https://leetcode.cn/problems/climbing-stairs/
/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function (n) {
  if (n === 1 || n === 2) return n;
  let i = n % 2 === 0 ? 0 : 1;
  let j = Math.trunc(n / 2); //看有多少个2
  let res = 0;
  while (j >= 0) {
    let tmp = 1;
    for (let k = i + j; k > j; k--) {
      tmp *= k / (k - j);
    }
    res += tmp;
    j--;
    i += 2;
  }
  return res;
};

// 优秀写法，动态规划
// 在第n台阶的方案是在n-1和n-2台阶方案的和
// var climbStairs = function (n) {
//   let p = 1,
//     q = 2,
//     res = 0;
//   if (n === 1) return p;
//   if (n === 2) return q;
//   for (let count = 3; count <= n; count++) {
//     res = p + q;
//     p = q;
//     q = res;
//   }
//   return res;
// };
// console.log(climbStairs(40));
