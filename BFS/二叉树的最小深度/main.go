package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := make([]*TreeNode, 0)
	step := 1
	q = append(q, root)
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			node := q[0]
			q = q[1:]
			if isLeaf(node) {
				return step
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		step++
	}
	return step
}

func isLeaf(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	return false
}
