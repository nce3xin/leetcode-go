package main

// 337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var memo = make(map[*TreeNode]int)

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if v, ok := memo[root]; ok {
		return v
	}
	// 抢root
	doRob := root.Val
	if root.Left != nil {
		doRob += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		doRob += rob(root.Right.Left) + rob(root.Right.Right)
	}
	// 不抢root
	doNotRob := rob(root.Left) + rob(root.Right)
	res := max(doRob, doNotRob)
	memo[root] = res
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
