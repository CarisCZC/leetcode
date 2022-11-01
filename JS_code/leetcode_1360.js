// leetcode 1360 日期间隔
/**
 * @param {string} date1
 * @param {string} date2
 * @return {number}
 */
var daysBetweenDates = function (date1, date2) {
  const dateMinllions1 = new Date(date1);
  const dateMinllions2 = new Date(date2);
  let days = 0;
  if (dateMinllions1 > dateMinllions2) {
    days = (dateMinllions1 - dateMinllions2) / (24 * 60 * 60 * 1000);
  } else {
    days = (dateMinllions2 - dateMinllions1) / (24 * 60 * 60 * 1000);
  }
  return days;
};

// 没啥本质区别，感觉只是服务器时间波动
// 优秀写法
// var daysBetweenDates = function (date1, date2) {
//   const aDay = 60 * 60 * 1000 * 24;

//   const res = Math.abs(new Date(date1).getTime() - new Date(date2).getTime());
//   return res / aDay;
// };
