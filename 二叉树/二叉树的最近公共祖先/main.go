package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, p, q *TreeNode) *TreeNode {
	return find(root, p.Val, q.Val)
}

// 在以root为根的二叉树中，寻找值为val1或者val2的节点。如果找到，返回此节点。找不到，返回nil
// 为什么要写这样一个奇怪的findNode函数呢？因为最近公共祖先系列问题的解法都是把这个函数作为框架的
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

// 找最近公共父节点，只要在findNode基础上，在后序位置添加一个判断逻辑,即可改造成寻找最近公共祖先的解法代码
// 返回的是最近公共祖先
func find(root *TreeNode, val1, val2 int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val1 || root.Val == val2 {
		return root
	}
	left := find(root.Left, val1, val2)
	right := find(root.Right, val1, val2)
	// 后序位置新增一个判断逻辑
	if left != nil && right != nil {
		return root
	}
	// 由于是后序遍历，自底向上，所以可以返回left就是最近公共父节点
	// 根节点不是p或者q。p、q不是一个在左子树，一个在右子树。说明p、q都在左子树，或者都在右子树。
	if left != nil {
		// 左子树找到了p或者q。上面分析，p、q这时候一定在同一个子树，所以p,q都在左子树。
		// 又是自底向上遍历，所以left就是最近公共父节点
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
