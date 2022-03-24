# 打家劫舍III

337 题。

第三题又想法设法地变花样了，此强盗发现现在面对的房子不是一排，不是一圈，而是一棵二叉树！房子在二叉树的节点上，相连的两个房子不能同时被抢劫。

整体的思路完全没变，还是做抢或者不抢的选择，取收益较大的选择。甚至我们可以直接按这个套路写出代码：

```
package main

// 337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var memo = make(map[*TreeNode]int)

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if v, ok := memo[root]; ok {
		return v
	}
	// 抢root
	doRob := root.Val
	if root.Left != nil {
		doRob += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		doRob += rob(root.Right.Left) + rob(root.Right.Right)
	}
	// 不抢root
	doNotRob := rob(root.Left) + rob(root.Right)
	res := max(doRob, doNotRob)
	memo[root] = res
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

```