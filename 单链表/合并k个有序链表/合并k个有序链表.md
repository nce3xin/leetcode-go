# 合并k个有序链表

合并 k 个有序链表的逻辑类似合并两个有序链表，难点在于，如何快速得到 k 个节点中的最小节点，接到结果链表上？

这里我们就要用到 优先级队列（二叉堆） 这种数据结构，把链表节点放入一个最小堆，就可以每次获得 k 个节点中的最小节点：

```java
ListNode mergeKLists(ListNode[] lists) {
    if (lists.length == 0) return null;
    // 虚拟头结点
    ListNode dummy = new ListNode(-1);
    ListNode p = dummy;
    // 优先级队列，最小堆
    PriorityQueue<ListNode> pq = new PriorityQueue<>(
        lists.length, (a, b)->(a.val - b.val));
    // 将 k 个链表的头结点加入最小堆
    for (ListNode head : lists) {
        if (head != null)
            pq.add(head);
    }

    while (!pq.isEmpty()) {
        // 获取最小节点，接到结果链表中
        ListNode node = pq.poll();
        p.next = node;
        if (node.next != null) {
            pq.add(node.next);
        }
        // p 指针不断前进
        p = p.next;
    }
    return dummy.next;
}
```

这个算法是面试常考题，它的时间复杂度是多少呢？

优先队列 pq 中的元素个数最多是 k，所以一次 poll 或者 add 方法的时间复杂度是 O(logk)；所有的链表节点都会被加入和弹出 pq，所以算法整体的时间复杂度是 O(Nlogk)，其中 k 是链表的条数，N 是这些链表的节点总数。

## 注意点

- Golang里要用heap，需要先引入container/heap，实现接口。如何知道该实现哪些接口呢？先import container/heap，然后随便在哪个地方调用一下heap.Push，ctrl+鼠标左键单击heap, 进入built-in实现，看到：
    ```
    type Interface interface {
        sort.Interface
        Push(x interface{}) // add x as element Len()
        Pop() interface{}   // remove and return element Len() - 1.
    }
    ```
  可以看出，这个堆结构继承自sort.Interface, 而sort.Interface，需要实现三个方法：Len() int /   Less(i, j int) bool  /  Swap(i, j int) 再加上堆接口定义的两个方法：Push(x interface{})   /  Pop() interface{}。故只要实现了这五个方法，变定义了一个堆。

  Less函数正常实现，是最小堆。反着实现，是最大堆。

  注意接口的Push和Pop方法是供heap包调用的，请使用heap.Push和heap.Pop来向一个堆添加或者删除元素。不要调用具体实现堆的Push和Pop方法。
