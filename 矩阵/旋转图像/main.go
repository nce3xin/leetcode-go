package main

// 48
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		reverseRow(matrix, n, i)
	}
}

func reverseRow(matrix [][]int, n int, row int) {
	p := 0
	q := n - 1
	for p < q {
		matrix[row][p], matrix[row][q] = matrix[row][q], matrix[row][p]
		p++
		q--
	}
}
