package main

import "fmt"

// 239

func main() {
	nums := []int{7, 2, 4}
	k := 2
	res := maxSlidingWindow(nums, k)
	fmt.Printf("res: %v\n", res)
}

type MonotonicQueue struct {
	q []int
}

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	n := len(nums)
	mq := newMonotonicQueue()
	// 先将前k个元素放入队列
	for i := 0; i < k; i++ {
		mq.push(nums[i])
	}
	// 记录前k个元素的最大值
	res = append(res, mq.max())
	for i := k; i < n; i++ {
		// 滑动窗口移除最前面的元素
		// 这里为什么是i-k，而不是i-k+1呢？其实可以用特例来证明。
		// 当i==k时候，我们需要删除的是nums[0], i-k 恰好是0.
		mq.pop(nums[i-k])
		// 滑动窗口添加最后面的元素
		mq.push(nums[i])
		// 记录最大值
		maxVal := mq.max()
		res = append(res, maxVal)
	}
	return res
}

func newMonotonicQueue() *MonotonicQueue {
	mq := &MonotonicQueue{q: []int{}}
	return mq
}

func (mq *MonotonicQueue) push(x int) {
	for !mq.isEmpty() && mq.back() < x {
		mq.q = mq.q[:len(mq.q)-1]
	}
	mq.q = append(mq.q, x)
}

func (mq *MonotonicQueue) pop(x int) {
	if mq.isEmpty() {
		return
	}
	if mq.front() == x {
		mq.q = mq.q[1:]
	}
}

func (mq *MonotonicQueue) max() int {
	return mq.front()
}

func (mq *MonotonicQueue) front() int {
	return mq.q[0]
}

func (mq *MonotonicQueue) back() int {
	return mq.q[len(mq.q)-1]
}

func (mq *MonotonicQueue) isEmpty() bool {
	if len(mq.q) == 0 {
		return true
	}
	return false
}
