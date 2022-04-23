package main

// 386
func lexicalOrder(n int) []int {
	res := make([]int, 0)
	for i := 1; i <= 9; i++ {
		dfs(n, &res, i)
	}
	return res
}

func dfs(n int, res *[]int, cur int) {
	// 就算有剪枝，这句if也不可省略。因为这题的dfs是从多个起点开始的。
	if cur > n {
		return
	}
	*res = append(*res, cur)
	for i := 0; i <= 9; i++ {
		num := cur*10 + i
		// 剪枝
		if num > n {
			continue
		}
		dfs(n, res, num)
	}
}
