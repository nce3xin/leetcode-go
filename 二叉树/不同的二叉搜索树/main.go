package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numTrees(n int) int {
	memo := make([][]int, n+1)
	for i, _ := range memo {
		memo[i] = make([]int, n+1)
	}
	return count(1, n, &memo)
}

// 定义：闭区间 [lo, hi] 的数字能组成 count(lo, hi) 种 BST
func count(lo, hi int, memo *[][]int) int {
	if lo > hi {
		return 1
	}
	if (*memo)[lo][hi] != 0 {
		return (*memo)[lo][hi]
	}
	res := 0
	for i := lo; i <= hi; i++ {
		leftNum := count(lo, i-1, memo)
		rightNum := count(i+1, hi, memo)
		res += leftNum * rightNum
	}
	(*memo)[lo][hi] = res
	return res
}
