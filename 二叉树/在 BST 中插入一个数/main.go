package main

// 701

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 找到空位置插入新节点
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}
