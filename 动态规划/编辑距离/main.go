package main

// 72

func minDistance(word1 string, word2 string) int {
	// dp[i][j]: s1[0...i-1]转化为s2[0...j-1]的最小编辑距离
	n1 := len(word1)
	n2 := len(word2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}
	// base case
	// dp[0][j]: 全部insert，j
	// dp[i][0]: 全部删除，i
	for i := 0; i <= n1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n2; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i][j-1]+1, // insert
					min(
						dp[i-1][j-1]+1, // replace
						dp[i-1][j]+1,   // delete
					),
				)
			}
		}
	}
	return dp[n1][n2]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
