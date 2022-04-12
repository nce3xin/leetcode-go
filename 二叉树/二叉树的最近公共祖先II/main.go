package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var foundP = false
var foundQ = false

func lowestCommonAncestor(root *TreeNode, p, q *TreeNode) *TreeNode {
	if (!foundP) || (!foundQ) {
		return nil
	}
	return find(root, p.Val, q.Val)
}

func findNode(root *TreeNode, val1, val2 int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val1 || root.Val == val2 {
		return root
	}
	left := findNode(root.Left, val1, val2)
	right := findNode(root.Right, val1, val2)
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

func find(root *TreeNode, val1, val2 int) *TreeNode {
	if root == nil {
		return nil
	}
	left := find(root.Left, val1, val2)
	right := find(root.Right, val1, val2)
	if left != nil && right != nil {
		return root
	}
	// 后序位置，判断当前节点是不是目标值
	// 想想最简单的后序遍历，这个地方可以把每个节点都判断一遍
	if root.Val == val1 || root.Val == val2 {
		if root.Val == val1 {
			foundP = true
		}
		if root.Val == val2 {
			foundQ = true
		}
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
