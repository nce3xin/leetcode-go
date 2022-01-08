package main

import "fmt"

func main() {
	//deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	deadends := []string{"0000"}
	//target := "0202"
	target := "8888"
	res := openLock(deadends, target)
	fmt.Printf("res: %v\n", res)
}

func addDeadendsToVisited(visited *map[string]bool, deadends []string) {
	for _, s := range deadends {
		(*visited)[s] = true
	}
}

func inDeadends(s string, deadends []string) bool {
	for _, d := range deadends {
		if s == d {
			return true
		}
	}
	return false
}

func openLock(deadends []string, target string) int {
	q := make([]string, 0)
	initStr := "0000"
	if inDeadends(initStr, deadends) {
		return -1
	}
	q = append(q, initStr)
	visited := make(map[string]bool)
	visited[initStr] = true
	addDeadendsToVisited(&visited, deadends)
	step := 0
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			if cur == target {
				return step
			}
			for j := 0; j < 4; j++ {
				plusStr := plusOne(cur, j)
				minusStr := minusOne(cur, j)
				if _, ok := visited[plusStr]; !ok {
					q = append(q, plusStr)
					visited[plusStr] = true
				}
				if _, ok := visited[minusStr]; !ok {
					q = append(q, minusStr)
					visited[minusStr] = true
				}
			}
		}
		step++
	}
	return -1
}

// s[i]位向上拨1次
func plusOne(s string, i int) string {
	res := []byte(s)
	if s[i] == '9' {
		res[i] = '0'
	} else {
		res[i] = s[i] + 1
	}
	return string(res)
}

// // s[i]位向下拨1次
func minusOne(s string, i int) string {
	res := []byte(s)
	if s[i] == '0' {
		res[i] = '9'
	} else {
		res[i] = s[i] - 1
	}
	return string(res)
}
