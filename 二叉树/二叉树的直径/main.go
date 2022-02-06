// 543
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	maxDepth(root, &res)
	return res
}

func maxDepth(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left, diameter)
	RightDepth := maxDepth(root.Right, diameter)
	*diameter = max(*diameter, leftDepth+RightDepth)
	return max(leftDepth, RightDepth) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
