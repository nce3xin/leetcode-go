package main

// 207

// DFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(numCourses, prerequisites)
	visited := make([]bool, numCourses)
	onPath := make([]bool, numCourses)
	hasCycle := false
	// 注意图中并不是所有节点都相连，所以要用一个 for 循环将所有节点都作为起点调用一次 DFS 搜索算法
	for i := 0; i < numCourses; i++ {
		traverse(graph, i, &visited, &onPath, numCourses, &hasCycle)
		if hasCycle {
			return false
		}
	}
	return true
}

// 参数start代表遍历的起点，类似于二叉树遍历函数traverse里的root，是必须要有的
func traverse(graph [][]int, start int, visited *[]bool, onPath *[]bool, n int, hasCycle *bool) {
	// 这两个if的顺序不能反，如果反了，会直接返回，没有设置hasCycle，导致解答错误
	if (*onPath)[start] {
		*hasCycle = true
		return
	}
	if (*visited)[start] {
		return
	}
	(*visited)[start] = true
	(*onPath)[start] = true
	for _, v := range graph[start] {
		traverse(graph, v, visited, onPath, n, hasCycle)
	}
	// 注意，这里不需要(*visited)[start] = false. 这也是visited数组和onPath数组的主要区别
	(*onPath)[start] = false
}

// 邻接表
func buildGraph(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}
	for _, edge := range prerequisites {
		// 这里 from := edge[0], to := edge[1]也可以
		from := edge[1]
		to := edge[0]
		graph[from] = append(graph[from], to)
	}
	return graph
}
