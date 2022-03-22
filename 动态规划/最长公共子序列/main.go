package main

// 1143

func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j]: text1前i个元素和text2前j个元素的最长公共子序列的长度
	n1 := len(text1)
	n2 := len(text2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}
	// base case
	// dp[0][...]: 0
	// dp[...][0]: 0
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(
					dp[i-1][j],
					max(
						dp[i][j-1],
						dp[i-1][j-1],
					),
				)
			}
		}
	}
	return dp[n1][n2]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
