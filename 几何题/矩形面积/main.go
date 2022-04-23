package main

// 223
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	area1 := (ay2 - ay1) * (ax2 - ax1)
	area2 := (by2 - by1) * (bx2 - bx1)
	overlapHeight := max(0, min(ay2, by2)-max(ay1, by1))
	overlapWidth := max(0, min(ax2, bx2)-max(ax1, bx1))
	overlapArea := overlapWidth * overlapHeight
	return area1 + area2 - overlapArea
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
