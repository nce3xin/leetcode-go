# 下一个更大元素II

[https://leetcode-cn.com/problems/next-greater-element-ii/](https://leetcode-cn.com/problems/next-greater-element-ii/)

如何处理环形数组。

给定一个循环数组 nums（nums[nums.length - 1]的下一个元素是nums[0]），返回nums中每个元素的 下一个更大元素 。

数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

一般是通过 % 运算符求模（余数），来获得环形特效。

这个问题肯定还是要用单调栈的解题模板，但难点在于，比如输入是 [2,1,2,4,3]，对于最后一个元素 3，如何找到元素 4 作为 Next Greater Number。

对于这种需求，常用套路就是将数组长度翻倍：

这样，元素 3 就可以找到元素 4 作为 Next Greater Number 了，而且其他的元素都可以被正确地计算。

有了思路，最简单的实现方式当然可以把这个双倍长度的数组构造出来，然后套用算法模板。但是，我们可以不用构造新数组，而是利用循环数组的技巧来模拟数组长度翻倍的效果。

```
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	s := make([]int, 0)
	for i := 2*n - 1; i >= 0; i-- {
		x := nums[i%n]
		for len(s) != 0 && s[len(s)-1] <= x {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			res[i%n] = -1
		} else {
			res[i%n] = s[len(s)-1]
		}
		s = append(s, x)
	}
	return res
}
```

## 时间复杂度

- 时间复杂度：O(n)