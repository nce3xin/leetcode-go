package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	minVal := min(p.Val, q.Val)
	maxVal := max(p.Val, q.Val)
	res := find(root, minVal, maxVal)
	return res
}

// 寻找最近公共祖先
func find(root *TreeNode, minVal, maxVal int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > maxVal {
		return find(root.Left, minVal, maxVal)
	}
	if root.Val < minVal {
		return find(root.Right, minVal, maxVal)
	}
	return root
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
