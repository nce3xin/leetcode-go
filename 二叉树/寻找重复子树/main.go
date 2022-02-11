// 652

package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	memo := make(map[string]int)
	res := make([]*TreeNode, 0)
	serializeTree(root, &memo, &res)
	return res
}

func serializeTree(root *TreeNode, memo *map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return "#"
	}
	leftStr := serializeTree(root.Left, memo, res)
	rightStr := serializeTree(root.Right, memo, res)
	// 这里有个大坑，int转string，不能用string(root.Val)这种方式，必须用strconv，否则会提示解答错误
	subtree := leftStr + "," + rightStr + "," + strconv.Itoa(root.Val)
	if cnt, ok := (*memo)[subtree]; !ok {
		(*memo)[subtree] = 1
	} else {
		if cnt == 1 {
			*res = append(*res, root)
		}
		(*memo)[subtree] = cnt + 1
	}
	return subtree
}
