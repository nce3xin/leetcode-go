package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return build(1, n)
}

// 闭区间[lo,hi]组成的所有可能的BST
func build(lo, hi int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if lo > hi {
		res = append(res, nil)
		return res
	}
	for i := lo; i <= hi; i++ {
		leftNodes := build(lo, i-1)
		rightNodes := build(i+1, hi)
		for _, ln := range leftNodes {
			for _, rn := range rightNodes {
				// 注意，root的创建，必须在最内循环。不能在第一层循环内，否则解答错误，会有重复答案。
				root := &TreeNode{
					Val:   i,
					Left:  nil,
					Right: nil,
				}
				root.Left = ln
				root.Right = rn
				res = append(res, root)
			}
		}
	}
	return res
}
