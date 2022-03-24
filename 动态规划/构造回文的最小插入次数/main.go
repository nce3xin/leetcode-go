package main

// 1312

func minInsertions(s string) int {
	// dp[i][j]: s[i...j] (闭区间) 成为回文串的最少插入次数
	n := len(s)
	dp := make([][]int, n)
	// base case
	// dp[i][j]: 当 i==j 时，只有1个字符，不用插入，就是回文串。为0
	// dp 数组已经全部初始化为 0，base case 已初始化
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	// 从下向上遍历。
	// 遍历顺序是由状态转移方程决定的，由于dp[i][j] 和 dp[i+1][j]，dp[i][j-1]，dp[i+1][j-1] 三个状态有关,
	// 为了保证每次计算 dp[i][j] 时，这三个状态都已经被计算，我们一般选择从下向上，从左到右遍历 dp 数组
	for i := n - 2; i >= 0; i-- {
		// 从左向右遍历
		// 注意，这里j是从i+1开始计算。因为dp数组的定义，就说明i<=j
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j]+1, dp[i][j-1]+1)
			}
		}
	}
	return dp[0][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
