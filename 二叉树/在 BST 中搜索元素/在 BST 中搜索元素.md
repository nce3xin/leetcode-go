# 在 BST 中搜索元素

力扣第 700 题「二叉搜索树中的搜索」就是让你在 BST 中搜索值为 target 的节点。

如果是在一棵普通的二叉树中寻找，可以这样写代码：

```c++
TreeNode searchBST(TreeNode root, int target);
    if (root == null) return null;
    if (root.val == target) return root;
    // 当前节点没找到就递归地去左右子树寻找
    TreeNode left = searchBST(root.left, target);
    TreeNode right = searchBST(root.right, target);

    return left != null ? left : right;
}
```

这样写完全正确，但这段代码相当于穷举了所有节点，适用于所有普通二叉树。那么应该如何充分利用信息，把 BST 这个「左小右大」的特性用上？

很简单，其实不需要递归地搜索两边，类似二分查找思想，根据 target 和 root.val 的大小比较，就能排除一边。我们把上面的思路稍稍改动：

```c++
TreeNode searchBST(TreeNode root, int target) {
    if (root == null) {
        return null;
    }
    // 去左子树搜索
    if (root.val > target) {
        return searchBST(root.left, target);
    }
    // 去右子树搜索
    if (root.val < target) {
        return searchBST(root.right, target);
    }
    return root;
}
```
另一种写法：

```c++
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val == root.Val {
		return root
	}
	if val > root.Val {
		return searchBST(root.Right, val)
	}
	return searchBST(root.Left, val)
}
```

两种写法就是后序和先序的区别。