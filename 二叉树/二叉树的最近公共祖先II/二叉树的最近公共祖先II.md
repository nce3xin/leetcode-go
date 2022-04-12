# 二叉树的最近公共祖先II

需要注意的是，标准的最近公共祖先题目，都明确告诉我们两个节点必定存在于二叉树中，如果没有这个前提条件，就需要修改代码了。

比如力扣第 1644 题「二叉树的最近公共祖先 II」：

给你输入一棵不含重复值的二叉树的，以及两个节点p和q，如果p或q不存在于树中，则返回空指针，否则的话返回p和q的最近公共祖先节点。

这题比较奇葩，居然没有go语言版本。

## 解题步骤

1. 设置两个全局标记变量：foundP和foundQ。
2. 把前序判断root是否是LCA的逻辑，挪到后序。

## 详解

在解决标准的最近公共祖先问题时，我们在find函数的前序位置有这样一段代码：

```
// 前序位置
if (root.val == val1 || root.val == val2) {
    // 如果遇到目标值，直接返回
    return root;
}
```

我也进行了解释，因为p和q都存在于树中，所以这段代码恰好可以解决最近公共祖先的第二种情况。

但对于这道题来说，p和q不一定存在于树中，所以你不能遇到一个目标值就直接返回，而应该对二叉树进行完全搜索（遍历每一个节点），如果发现p或q不存在于树中，那么是不存在LCA的。

回想我在文章开头分析的几种find函数的写法，哪种写法能够对二叉树进行完全搜索来着？

这种：

```
TreeNode find(TreeNode root, int val) {
    if (root == null) {
        return null;
    }
    // 先去左右子树寻找
    TreeNode left = find(root.left, val);
    TreeNode right = find(root.right, val);
    // 后序位置，判断 root 是不是目标节点
    if (root.val == val) {
        return root;
    }
    // root 不是目标节点，再去看看哪边的子树找到了
    return left != null ? left : right;
}
```

那么解决这道题也是类似的，我们只需要把前序位置的判断逻辑放到后序位置即可：

```
// 用于记录 p 和 q 是否存在于二叉树中
boolean foundP = false, foundQ = false;

TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
    TreeNode res = find(root, p.val, q.val);
    if (!foundP || !foundQ) {
        return null;
    }
    // p 和 q 都存在二叉树中，才有公共祖先
    return res;
}

// 在二叉树中寻找 val1 和 val2 的最近公共祖先节点
TreeNode find(TreeNode root, int val1, int val2) {
    if (root == null) {
        return null;
    }
    TreeNode left = find(root.left, val1, val2);
    TreeNode right = find(root.right, val1, val2);

    // 后序位置，判断当前节点是不是 LCA 节点
    if (left != null && right != null) {
        return root;
    }

    // 后序位置，判断当前节点是不是目标值
    if (root.val == val1 || root.val == val2) {
        // 找到了，记录一下
        if (root.val == val1) foundP = true;
        if (root.val == val2) foundQ = true;
        return root;
    }

    return left != null ? left : right;
}
```

这样改造，对二叉树进行完全搜索，同时记录p和q是否同时存在树中，从而满足题目的要求。