package main

import "fmt"

// https://leetcode-cn.com/problems/n-queens/

func main() {
	n := 4
	res := solveNQueens(n)
	fmt.Printf("res: %v\n", res)
}

func solveNQueens(n int) [][]string {
	board := constructBoard(n)
	res := make([][]string, 0)
	backtrack(board, &res, 0, n)
	return res
}

func constructBoard(n int) [][]string {
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i] = append(board[i], ".")
		}
	}
	return board
}

func backtrack(board [][]string, res *[][]string, i, n int) {
	if i == n {
		s := constructString(board, n)
		*res = append(*res, s)
		return
	}
	for j := 0; j < n; j++ {
		if !isValid(board, i, j, n) {
			continue
		}
		board[i][j] = "Q"
		backtrack(board, res, i+1, n)
		board[i][j] = "."
	}
}

func constructString(s [][]string, n int) []string {
	res := make([]string, 0)
	for i := 0; i < n; i++ {
		tmp := ""
		for j := 0; j < n; j++ {
			tmp = tmp + s[i][j]
		}
		res = append(res, tmp)
	}
	return res
}

func isValid(board [][]string, i, j, n int) bool {
	// 左上角
	for x, y := i-1, j-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
		if board[x][y] == "Q" {
			return false
		}
	}
	// 上方
	for x := i - 1; x >= 0; x-- {
		if board[x][j] == "Q" {
			return false
		}
	}
	// 右上方
	for x, y := i-1, j+1; x >= 0 && y < n; x, y = x-1, y+1 {
		if board[x][y] == "Q" {
			return false
		}
	}
	return true
}
