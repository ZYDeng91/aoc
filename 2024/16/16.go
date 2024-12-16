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

type State struct {
	pos Loc
	dir Loc
}

type Node struct {
	pos Loc
	dir Loc
	score int
}

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
	content, err := os.ReadFile("16.in")
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")

	mat := make([]string, 0)

	for _, line := range lines {
		mat = append(mat, line)
	}

	fmt.Println(day16(mat))
}

func day16(mat []string) (int, int) {
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
	res, res2 := astar(start, end, mat)

	return res, res2
}

func astar(start, end Loc, mat []string) (int, int) {
	//openSet := &NodeHeap{Node{start, Loc{0,0}, 0}}
	// start at facing *East*
	openSet := &NodeHeap{Node{start, Loc{0,1}, 0}}
	heap.Init(openSet)
	gScore := make(map[State]int)
	//gScore[State{start, Loc{0,0}}] = 0
	gScore[State{start, Loc{0,1}}] = 0
	dirs := make([]Loc, 4)
	for i:=0;i<4;i++ {
		dirs[i] = Loc{(i-2)%2, (i-1)%2}
	}

	for openSet.Len() > 0 {
		currentNode := heap.Pop(openSet).(Node)
		current := currentNode.pos
		if current == end {
			return currentNode.score, backtrack(gScore, currentNode, mat)
		}
		for _, dir := range(dirs) {
			// no turning back
			if dir.x == -currentNode.dir.x && dir.y == -currentNode.dir.y {
				continue
			}
			next := Loc{current.x+dir.x, current.y+dir.y}
			if mat[next.x][next.y] == '#' {
				continue
			}
			tentative_g := currentNode.score
			if dir == currentNode.dir {
				tentative_g += 1
			} else {
				tentative_g += 1001
			}
			next_g, _ := gScore[State{next, dir}]
			if next_g == 0 || (tentative_g <= next_g) {
				gScore[State{next, dir}] = tentative_g
				heap.Push(openSet, Node{next, dir, tentative_g})
			}
			
		}
		
	}

	return 0, 0
}

func backtrack(gScore map[State]int, current Node, mat []string) int {
	visited := make(map[Loc]bool)
	visited[current.pos] = true
	queue := []Node{current}
	for len(queue) > 0 {
		new_queue := make([]Node, 0)
		for _, currentNode := range(queue) {
			for key, val := range(gScore) {
				if key.pos.x==currentNode.pos.x-currentNode.dir.x && key.pos.y==currentNode.pos.y-currentNode.dir.y {
					if (currentNode.dir == key.dir && currentNode.score == val+1) || (currentNode.dir != key.dir && currentNode.score == val+1001) {
						visited[key.pos] = true
						new_queue = append(new_queue, Node{key.pos, key.dir, val})
					}
				}
			}
		}
		queue = new_queue
	}

	res := 0

	for _ = range(visited) {
		res += 1
	}

	// debug display
	/*
	for x, row := range(mat) {
		for y, cell := range(row) {
			if visited[Loc{x, y}] {
				fmt.Print("O")
			} else if cell == '.' {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	*/
	return res
}
