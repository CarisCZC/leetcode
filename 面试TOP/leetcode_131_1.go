package main

func partition(s string) (ans [][]string) {
	n := len(s)
	f := make([][]bool, n)
	for i := range f { // 初始化状态矩阵
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			// f[i][j]表示，从i到j的字符串是一个回文串
			// i到j是回文串的要求是 s[i] == s[j]且i+1到j-1构成了回文串
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}
	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1]) //当前数组
				dfs(j + 1)
				// 回溯过程，将前面的加入的当前数组删除掉
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return
}

func main() {
	partition("aab")
}
