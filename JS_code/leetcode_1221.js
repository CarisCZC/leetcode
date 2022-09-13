// leetcode 1221. 分割平衡字符串
// https://leetcode.cn/problems/split-a-string-in-balanced-strings/
/**
 * @param {string} s
 * @return {number}
 */
// var balancedStringSplit = function (s) {
//   // 需要一个栈结构
//   const stack = [];
//   let res = 0;
//   for (const ch of s) {
//     if (stack.length === 0 || stack.at(-1) === ch) {
//       stack.push(ch);
//       continue;
//     }
//     stack.pop();
//     if (stack.length === 0) res++;
//   }
//   return res;
// };

var balancedStringSplit = function (s) {
  let res = 0,
    n = 0;
  for (const ch of s) {
    if (ch === "L") n++;
    else n--;
    if (n === 0) res++;
  }
  return res;
};
