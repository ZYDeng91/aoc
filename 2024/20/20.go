package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type Loc struct {
	x int
	y int
}

type Node struct {
	pos Loc
	score int
}

var gScore map[Loc]int

type NodeHeap []Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].score < h[j].score }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Node))
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getIndex(h *NodeHeap, target Loc) int {
	for i, val := range *h {
		if val.pos == target {
			return i
		}
	}
	return -1
}

func main() {
	content, err := os.ReadFile("20.in")
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")

	mat := make([][]byte, 0)

	for _, line := range lines {
		mat = append(mat, []byte(line))
	}

	fmt.Println(day20(mat, 100, 20))
}

func day20(mat [][]byte, filter, cheats int) int {
	var start, end Loc
	for x, row := range(mat) {
		for y, cell := range(row) {
			if cell == 'S' {
				start = Loc{x, y}
			}
			if cell == 'E' {
				end = Loc{x, y}
			}
		}
	}
	//fmt.Println(start, end, mat)
	res := 0
	baseline := astar(start, end, mat)
	fmt.Println(baseline)
	/*for x, row := range(mat) {
		if x == 0 || x == 140 {
			continue
		}
		for y, cell := range(row) {
			if y == 0 || y == 140 {
				continue
			}
			if cell == '#' {
				temp := cp(mat)
				temp[x][y] = '.'
				if astar(start, end, temp) <= baseline - 100 {
					res += 1
				}
			}
		}
	}*/
	for k, v := range(gScore){
		for k2, v2 := range(gScore){
			delta := v2-v
			distance := abs(k2.x-k.x)+abs(k2.y-k.y)
			// only one direction to avoid repeats
			if delta - distance < filter || distance > cheats {
				continue
			}

			//fmt.Println(k, v, k2, v2, delta-distance)
			res += 1
		}
	}

	return res
}

func astar(start, end Loc, mat [][]byte) int {
	openSet := &NodeHeap{Node{start, 0}}
	heap.Init(openSet)
	//gScore := make(map[Loc]int)
	gScore = make(map[Loc]int)
	gScore[start] = 0
	dirs := make([]Loc, 4)
	for i:=0;i<4;i++ {
		dirs[i] = Loc{(i-2)%2, (i-1)%2}
	}

	for openSet.Len() > 0 {
		currentNode := heap.Pop(openSet).(Node)
		current := currentNode.pos
		if current == end {
			return currentNode.score
		}
		for _, dir := range(dirs) {
			next := Loc{current.x+dir.x, current.y+dir.y}
			if mat[next.x][next.y] == '#' {
				continue
			}
			tentative_g := currentNode.score + 1
			next_g, ok := gScore[next]
			if !ok || tentative_g <= next_g {
				gScore[next] = tentative_g
				heap.Push(openSet, Node{next, tentative_g})
			}
			
		}
		
	}

	return 0
}

func cp(mat [][]byte) [][]byte {
	res := make([][]byte, len(mat))
	for x, row := range(mat) {
		newrow := make([]byte, len(mat[0]))
		for y, cell := range(row) {
			newrow[y] = cell
		}
		res[x] = newrow
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
