package main

// 22

// 题目改写：现在有 2n 个位置，每个位置可以放置字符 ( 或者 )，组成的所有括号组合中，有多少个是合法的？
func generateParenthesis(n int) []string {
	track := make([]byte, 0)
	res := make([]string, 0)
	backtrack(&track, &res, n, n)
	return res
}

// left: 剩下可用的左括号数量
// right：剩下可用的右括号数量
func backtrack(track *[]byte, res *[]string, left, right int) {
	// 可用的左括号更多，说明已经用掉的右括号更多，不合法
	if right < left {
		return
	}
	if left < 0 || right < 0 {
		return
	}
	if left == 0 && right == 0 {
		*res = append(*res, string(*track))
		return
	}
	*track = append(*track, '(')
	backtrack(track, res, left-1, right)
	*track = (*track)[:len(*track)-1]

	*track = append(*track, ')')
	backtrack(track, res, left, right-1)
	*track = (*track)[:len(*track)-1]
}
