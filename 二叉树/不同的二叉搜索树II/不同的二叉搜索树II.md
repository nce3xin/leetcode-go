# 不同的二叉搜索树II

那么，如果给一个进阶题目，不止让你计算有几个不同的 BST，而是要你构建出所有合法的 BST，如何实现这个算法呢？

这道题就是力扣第 95 题「不同的二叉搜索树 II」，让你构建所有 BST。

明白了上道题构造合法 BST 的方法，这道题的思路也是一样的：

1、穷举root节点的所有可能。

2、递归构造出左右子树的所有合法 BST。

3、给root节点穷举所有左右子树的组合。

其实就是后序遍历。

```golang
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

```