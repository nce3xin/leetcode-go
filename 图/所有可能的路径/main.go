package main

// 797

func allPathsSourceTarget(graph [][]int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	n := len(graph)
	traverse(graph, 0, &res, &path, n)
	return res
}

func traverse(graph [][]int, start int, res *[][]int, path *[]int, n int) {
	// 这句必须在最开始，不能在if语句后面，否则path会少一个元素
	*path = append(*path, start)
	// 到达终点
	if start == n-1 {
		t := make([]int, len(*path))
		copy(t, *path)
		*res = append(*res, t)
		// 可以在这直接 return，但要 removeLast 正确维护 path
		// *path = (*path)[:len(*path)-1];
		// return;
		// 不 return 也可以，因为图中不包含环，不会出现无限递归
	}
	// 递归每个相邻节点
	for _, v := range graph[start] {
		traverse(graph, v, res, path, n)
	}
	// 从路径中移出节点start，回溯
	*path = (*path)[:len(*path)-1]
}
