// https://leetcode.cn/problems/available-captures-for-rook/
/**
 * @param {character[][]} board
 * @return {number}
 */
var numRookCaptures = function (board) {
  //先找车,记录车的坐标，同时记下其他棋子的坐标
  let index = [0, 0];
  let others = [];
  for (let i = 0; i < 8; i++) {
    for (let j = 0; j < 8; j++) {
      if (board[i][j] === "R") {
        index = [i, j];
      } else if (board[i][j] !== ".") {
        others.push([board[i][j], i, j]);
      }
    }
  }
  // 找目标棋子
  let res = 0;
  // 四个位置，上下左右
  let target = new Array(4).fill(["", 0, 0]);
  for (const other of others) {
    // 在一列上，即上下
    if (other[2] === index[1]) {
      // 上
      if (index[0] > other[1]) {
        target[0] =
          !target[0][0] || other[1] > target[0][1] ? other : target[0];
      } else {
        // 下
        target[1] =
          !target[1][0] || other[1] < target[1][1] ? other : target[1];
      }
    }
    // 在一行上，即左右
    if (other[1] === index[0]) {
      if (other[2] < index[1]) {
        // 左
        target[2] =
          !target[2][0] || other[2] > target[2][2] ? other : target[2];
      }
      if (other[2] > index[1]) {
        target[3] =
          !target[3][0] || other[2] < target[3][2] ? other : target[3];
      }
    }
  }
  for (const tar of target) {
    if (tar[0] === "p") res++;
  }
  return res;
};
