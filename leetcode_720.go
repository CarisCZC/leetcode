//leetcode 720,词典中最长的单词
package main

// map解法
// func longestWord(words []string) (longest string) {
// 	sort.Slice(words, func(a, b int) bool {
// 		s, t := words[a], words[b]
// 		return len(s) < len(t) || len(s) == len(t) && s > t
// 	})

// 	candidates := map[string]struct{}{"": {}}
// 	for _, word := range words {
// 		if _, ok := candidates[word[:len(word)-1]]; ok {
// 			longest = word
// 			candidates[word] = struct{}{}
// 		}
// 	}
// 	return longest
// }

// 字典树tire（前缀树）
type Trie struct {
	children [26]*Trie //每个位置存的是一个字符
	isEnd    bool      //end并不是表示是否是一个叶节点，也是表示是否是一个单词结尾。
	// 比如，我在一个树里存了 "ap" 和"app"，"ap"是由子节点的，即"app"，但"ap","app"的isEnd都是true
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil || !node.children[ch].isEnd {
			//isEnd = false，说明没有这个单词
			// 在这个题中，要求一个路径上的所有isEnd都是true
			return false
		}
		node = node.children[ch]
	}
	return true
}

func longestWord(words []string) (longest string) {
	t := &Trie{}
	for _, word := range words {
		t.Insert(word)
	}
	for _, word := range words {
		if t.Search(word) && (len(word) > len(longest) || len(word) == len(longest) && word < longest) {
			longest = word
		}
	}
	return longest

}
