// leetcode 997 小镇法官
/**
 * @param {number} n
 * @param {number[][]} trust
 * @return {number}
 */
// 编号从1-n 数组。
// 两个n数组，第一个n数组表示，不信任其他人的人。
// 第二个数组表示某个人被多少人信任
var findJudge = function (n, trust) {
  const trustOne = [];
  const beTrusted = [];
  // 初始化数组
  for (let i = 0; i < n; i++) {
    trustOne.push(0);
    beTrusted.push(0);
  }
  // 遍历trust统计
  trust.forEach((v) => {
    trustOne[v[0] - 1]++;
    beTrusted[v[1] - 1]++;
  });
  let res = -1;
  trustOne.forEach((v, i) => {
    if (v === 0 && beTrusted[i] === n - 1) {
      res = i + 1;
    }
  });
  // const res = trustOne.find((v, i) => {
  //   if (v === 0) return i;
  // });
  // if (res === -1 || beTrusted[res] !== n - 1) return -1;
  return res;
};

const n = 3,
  trust = [
    [1, 3],
    [2, 3],
  ];

console.log(findJudge(n, trust));
