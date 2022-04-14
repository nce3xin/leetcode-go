# 搜索二维矩阵II

[https://leetcode-cn.com/problems/search-a-2d-matrix-ii/](https://leetcode-cn.com/problems/search-a-2d-matrix-ii/)

编写一个高效的算法来搜索二维矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

- 每行的元素从左到右升序排列。
- 每列的元素从上到下升序排列。

## 解法

从右上角开始搜索。

如果 target 的值小于当前值，也就意味着当前值所在的列肯定不会存在 target 了，可以把当前列去掉，从新的右上角的值开始遍历。

同理，如果 target 的值大于当前值，也就意味着当前值所在的行肯定不会存在 target 了，可以把当前行去掉，从新的右上角的值开始遍历。

在搜索的过程中，如果超出了矩阵的边界，那么说明矩阵中不存在 target。

## 时间复杂度

O(m+n)。

每次干掉一行或者一列，所以是O(m+n)。

