// leecode 154. 一年中的第几天
// https://leetcode.cn/problems/day-of-the-year/
/**
 * @param {string} date
 * @return {number}
 */
var dayOfYear = function (date) {
  // 先获得月
  // 每个月的天数，二月先空着
  let res = +date.slice(8, date.length);

  const months = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
  const year = +date.slice(0, 4);
  const month = +date.slice(5, 7);
  console.log(year);
  console.log(month);
  console.log(res);
  months.forEach((v, i) => {
    if (i >= month - 1) return;
    res += v;
    if (i == 1) {
      if (year % 400 === 0 || (year % 100 !== 0 && year % 4 === 0)) res += 1;
    }
  });
  return res;
};

console.log(dayOfYear("2019-02-9"));
