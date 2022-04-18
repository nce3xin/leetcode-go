# Excel列名称

[https://leetcode-cn.com/problems/excel-sheet-column-title/](https://leetcode-cn.com/problems/excel-sheet-column-title/)

## 解法

26进制转换。

这是一道从 1 开始的的 26 进制转换题。

对于一般性的进制转换题目，只需要不断地对 columnNumber 进行 % 运算取得最后一位，然后对 columnNumber 进行 / 运算，直到 columnNumber 为 0 即可。

但本题是从1开始的。正常进制转换，是从0开始的。

因此在执行「进制转换」操作前，我们需要先对 columnNumber 执行减一操作，从而实现整体偏移。

也就是在for循环内，columnNumber 先减 1。
