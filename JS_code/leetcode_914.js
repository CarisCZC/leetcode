// leetcode 914. 卡牌分组
// https://leetcode.cn/problems/x-of-a-kind-in-a-deck-of-cards/
/**
 * @param {number[]} deck
 * @return {boolean}
 */
// 出现的次数有公因数
var hasGroupsSizeX = function (deck) {
  const map = new Map();
  deck.forEach((value) => {
    if (!map.get(`${value}`)) map.set(`${value}`, 0);
    map.set(`${value}`, map.get(`${value}`) + 1);
  });
  let res = 0;
  map.forEach((value) => {
    if (res === 0) res = value;
    else {
      res = maxcommonDivisor(value, res);
    }
  });
  return !(res < 2);
};

const maxcommonDivisor = function (a, b) {
  if (a < b) {
    [a, b] = [b, a];
  }
  if (a % b === 0) return b;
  return maxcommonDivisor(b, a % b);
};

console.log(hasGroupsSizeX([1, 1, 1, 1, 2, 2, 2, 2, 2, 2]));
