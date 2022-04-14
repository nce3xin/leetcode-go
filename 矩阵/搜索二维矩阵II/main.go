package main

// 240

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	i := 0
	j := n - 1
	for i >= 0 && i < m && j >= 0 && j < n {
		v := matrix[i][j]
		if v < target {
			i++
		} else if v > target {
			j--
		} else {
			return true
		}
	}
	return false
}
