// leetcode_1160 1160. 拼写单词
//https://leetcode.cn/problems/find-words-that-can-be-formed-by-characters/
/**
 * @param {string[]} words
 * @param {string} chars
 * @return {number}
 */

// 在words的每个单词循环chars中重复的
var countCharacters = function (words, chars) {
  let res = 0;
  const charMap = new Map();
  for (let i = 0; i < chars.length; i++) {
    if (!charMap.has(chars[i])) charMap.set(chars[i], 0);
    charMap.set(chars[i], charMap.get(chars[i]) + 1);
  }
  words.forEach((value) => {
    const tmpMap = new Map(charMap);
    for (let i = 0; i < value.length; i++) {
      if (!tmpMap.has(value[i])) return;
      const tmpV = tmpMap.set(value[i], tmpMap.get(value[i]) - 1).get(value[i]);
      if (tmpV < 0) return;
    }
    res += value.length;
  });
  return res;
};

console.log(countCharacters(["cat", "bt", "hat", "tree"], "atach"));
