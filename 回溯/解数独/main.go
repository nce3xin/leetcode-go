package main

// 37
func solveSudoku(board [][]byte) {
	n := len(board)
	backtrack(board, n, 0, 0)
}

// 这儿有个坑，backtrack函数必须有布尔类型的返回值，否则会填不上数字。
func backtrack(board [][]byte, n int, i, j int) bool {
	// 穷举到最后一列的话就换到下一行重新开始。
	if j == n {
		return backtrack(board, n, i+1, 0)
	}
	if i == n {
		// 找到一个可行解，触发 base case
		return true
	}
	if board[i][j] != '.' {
		// 如果有预设数字，不用我们穷举
		return backtrack(board, n, i, j+1)
	}
	choices := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, c := range choices {
		// 如果遇到不合法的数字，就跳过
		if !isValid(board, n, i, j, c) {
			continue
		}
		board[i][j] = c
		// 如果找到一个可行解，立即结束
		if backtrack(board, n, i, j+1) {
			return true
		}
		board[i][j] = '.'
	}
	// 穷举完 1~9，依然没有找到可行解，此路不通
	return false
}

func isValid(board [][]byte, n int, row, col int, curChoice byte) bool {
	// 判断当前行是否存在重复项
	for j := 0; j < n; j++ {
		if board[row][j] == curChoice {
			return false
		}
	}
	// 判断当前列是否存在重复
	for i := 0; i < n; i++ {
		if board[i][col] == curChoice {
			return false
		}
	}
	// 判断当前3*3的方框内，是否存在重复项
	r := row / 3
	c := col / 3
	for i := 3 * r; i < 3*r+3; i++ {
		for j := 3 * c; j < 3*c+3; j++ {
			if board[i][j] == curChoice {
				return false
			}
		}
	}
	return true
}
