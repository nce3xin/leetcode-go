package main

// 450

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		// delete
		// 叶子节点
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// 只有右节点
		if root.Left == nil {
			return root.Right
		}
		// 只有左节点
		if root.Right == nil {
			return root.Left
		}
		// 左右子树都有，用右子树的最小值替换根节点
		min := findMin(root.Right)
		root.Right = deleteNode(root.Right, min.Val)
		min.Left = root.Left
		min.Right = root.Right
		return min
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}

func findMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}
