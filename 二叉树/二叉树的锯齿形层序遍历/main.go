package main

// 103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	level := 1
	for len(q) > 0 {
		size := len(q)
		tmp := make([]int, size)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
			if level%2 != 0 {
				// 尾插
				tmp[i] = cur.Val
			} else {
				// 头插
				tmp[size-i-1] = cur.Val
			}
		}
		res = append(res, tmp)
		level++
	}
	return res
}
