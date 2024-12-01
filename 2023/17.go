package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	x int
	y int
}
type Node struct {
	loc   Coord
	dir   Coord
	score int
}
type Move struct {
	loc Coord
	dir Coord
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

func getIndex(h *NodeHeap, target Coord) int {
	for i, val := range *h {
		if val.loc == target {
			return i
		}
	}
	return -1
}

func main() {
	content, err := os.ReadFile("17.in")
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")

	mat := make([][]int, len(lines))

	for i, line := range lines {
		newline := make([]int, len(line))
		for j, item := range line {
			newline[j] = int(item - '0')
		}
		mat[i] = newline
	}

	mat = mat[:len(mat)-1][:]

	fmt.Println(day17(mat))
}

func day17(mat [][]int) int {
	start := Coord{0, 0}
	goal := Coord{len(mat) - 1, len(mat[0]) - 1}

	h := func(input Coord) int {
		//return 5*(goal.x-input.x+goal.y-input.y)
		return 0
	}

	return astar(start, goal, mat, h)
}

func astar(start, goal Coord, mat [][]int, h func(Coord) int) int {
	openSet := &NodeHeap{Node{start, Coord{-1, -1}, 0}}
	heap.Init(openSet)
	gScore := make(map[Move]int)
	gScore[Move{start, Coord{0, 0}}] = 0
	min := 4
	max := 10

	dirs := []Coord{Coord{0, 1}, Coord{0, -1}, Coord{1, 0}, Coord{-1, 0}}
	for openSet.Len() > 0 {
		current1 := heap.Pop(openSet).(Node)
		current := current1.loc
		if current == goal {
			return current1.score
		}
		for _, dir := range dirs {
			if dir.x == current1.dir.x {
				continue
			}
			if dir.y == current1.dir.y {
				continue
			}
			tentative_g := current1.score
			for i := 1; i <= max; i++ {
				neighbor := Coord{current.x + i*dir.x, current.y + i*dir.y}
				if neighbor.x < 0 || neighbor.y < 0 || neighbor.x > goal.x || neighbor.y > goal.y {
					continue
				}
				tentative_g += mat[neighbor.x][neighbor.y]
				if i < min {
					continue
				}
				neighbor_g, _ := gScore[Move{neighbor, dir}]
				if neighbor_g == 0 || (tentative_g < neighbor_g) {
					gScore[Move{neighbor, dir}] = tentative_g
					heap.Push(openSet, Node{neighbor, dir, tentative_g})
				}
			}
		}
	}

	return -1
}
