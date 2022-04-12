package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度：O(N)，N二叉树中的节点个数
// 空间复杂度：O(N)，N二叉树中的节点个数
// 空间复杂度主要取决于递归调用层数，最大层数等于二叉树的高度，最坏情况下，二叉树的高度等于二叉树中的节点个数
func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	traverse(root, &res)
	// 这里返回的是res，而不是traverse的返回值
	return res
}

// 返回以root为起点的最大单边路径和
func traverse(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	// 这里有个坑，需要和0取最大值。因为，如果是负值，还不如不取
	leftMaxSum := max(0, traverse(root.Left, res))
	rightMaxSum := max(0, traverse(root.Right, res))
	maxSum := max(leftMaxSum, rightMaxSum) + root.Val
	// 注意，这里更新res，不是用maxSum，而是 leftMaxSum + rightMaxSum + root.Val
	// 要想在每个节点遍历的时候，都更新res，就是在这儿更新的，你想想，前中后序遍历，输出所有节点的值，不就是这样输出的么
	*res = max(*res, leftMaxSum+rightMaxSum+root.Val)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
