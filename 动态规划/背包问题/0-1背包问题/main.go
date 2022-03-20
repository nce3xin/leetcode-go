package main

import "fmt"

func main() {
	m := 10
	A := []int{2, 3, 5, 7}
	V := []int{1, 5, 2, 4}
	res := backPackII(m, A, V)
	// expected: 9
	fmt.Printf("res: %v\n", res)
}

// m: 背包的可装载重量
// A: 各个物品的重量
// V: 各个物品的价值
func backPackII(m int, A []int, V []int) int {
	n := len(A) // 物品个数
	// dp[i][w] 的定义如下：对于前 i 个物品，当前背包的剩余容量为 w，这种情况下可以装的最大价值是 dp[i][w]。
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ { // 注意这里要从0开始初始化
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if j >= A[i-1] {
				// 最后为什么是A[i-1]和V[i-1]，而不是A[i]和V[i]?不是应该是第i件物品的重量和价值么？
				// 没错呀！就应该是第i件物品的重量和价值，但请看好，函数参数的A和V数组，是从0开始的呀！
				// A和V数组，不是dp数组！
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-A[i-1]]+V[i-1])
			} else { // 这个else分支不可删除，和零钱兑换那题不一样
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][m]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
