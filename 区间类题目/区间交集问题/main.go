package main

// 986

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := make([][]int, 0)
	i := 0
	j := 0
	n1 := len(firstList)
	n2 := len(secondList)
	for i < n1 && j < n2 {
		x1 := firstList[i][0]
		y1 := firstList[i][1]
		x2 := secondList[j][0]
		y2 := secondList[j][1]
		// 没有交集的条件是：y1 < x2 || x1 > y2
		// 根据否命题的定义，有交集的条件就是：y1 >= x2 && x1 <= y2
		if y1 >= x2 && x1 <= y2 {
			overlapY := min(y1, y2)
			overlapX := max(x1, x2)
			res = append(res, []int{overlapX, overlapY})
		}
		if y1 > y2 {
			j++
		} else {
			i++
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
