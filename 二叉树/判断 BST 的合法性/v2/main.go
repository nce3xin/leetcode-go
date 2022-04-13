package main

// 98

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	res := isBST(root, nil, nil)
	return res
}

func isBST(root *TreeNode, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	// 这里有个坑，不能写成：isLeftBST := isBST(root.Left, null, root)
	isLeftBST := isBST(root.Left, min, root)
	isRightBST := isBST(root.Right, root, max)
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	if isLeftBST && isRightBST {
		return true
	}
	return false
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
