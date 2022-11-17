// https://leetcode.cn/problems/matrix-cells-in-distance-order/
/**
 * @param {number} rows
 * @param {number} cols
 * @param {number} rCenter
 * @param {number} cCenter
 * @return {number[][]}
 */
var allCellsDistOrder = function (rows, cols, rCenter, cCenter) {
  // 按照不同的距离分坐标
  const distHash = new Array(cols + rows - 1).fill(0).map((item) => []);
  for (let i = 0; i < rows; i++) {
    for (let j = 0; j < cols; j++) {
      distHash[dist(rCenter, cCenter, i, j)].push([i, j]);
    }
  }
  // 扁平化
  return distHash.flat(1);
};
const dist = function (rCenter, cCenter, curRow, curCol) {
  return Math.abs(rCenter - curRow) + Math.abs(cCenter - curCol);
};

console.log(allCellsDistOrder(2, 3, 1, 2));
