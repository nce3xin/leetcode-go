package main

// 110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	res := true
	getHeight(root, &res)
	return res
}

func getHeight(root *TreeNode, isBalanced *bool) int {
	if root == nil {
		return 0
	}
	leftHeight := getHeight(root.Left, isBalanced)
	rightHeight := getHeight(root.Right, isBalanced)
	if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
		*isBalanced = false
	}
	height := max(leftHeight, rightHeight) + 1
	return height
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
