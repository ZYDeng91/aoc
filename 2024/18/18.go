package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Loc struct {
	x int
	y int
}
type State struct {
	pos Loc
	step int
}
type Node struct {
	pos Loc
	step int
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
	content, err := os.Open("18.go")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	size := 71

	mat := make([][]byte, size)

	for i, _ := range(mat) {
		newrow := make([]byte, size)
		for j, _ := range(newrow) {
			newrow[j] = '.'
		}
		mat[i] = newrow
	}

	blocks := make([]Loc, 0)

	for scanner.Scan() {
		text := scanner.Text()
		nums_str := strings.Split(text, ",")
		temp1, _ := strconv.Atoi(nums_str[1])
		temp2, _ := strconv.Atoi(nums_str[0])
		blocks = append(blocks, Loc{temp1, temp2})
	}

	fmt.Println(day18(mat, blocks[:1024], size))
	//fmt.Println(day18(mat, blocks, size))
}

func day18(mat [][]byte, blocks []Loc, size int) int {
	start := Loc{0, 0}
	end := Loc{size-1, size-1}
	for _, block := range(blocks) {
		mat[block.x][block.y] = '#'
	}
	//display(mat)
	res := astar(start, end, mat)

	return res
}

func astar(start, end Loc, mat [][]byte) int {
	//openSet := &NodeHeap{Node{start, Loc{0,0}, 0}}
	// start at facing *East*
	openSet := &NodeHeap{Node{start, 0, 0}}
	heap.Init(openSet)
	gScore := make(map[State]int)
	//gScore[State{start, Loc{0,0}}] = 0
	gScore[State{start, 0}] = 0
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
			if next.x < 0 || next.y < 0 || next.x > end.x || next.y > end.y {
				continue
			}
			if mat[next.x][next.y] == '#' {
				continue
			}
			tentative_g := currentNode.score
			tentative_g += 1
			next_g, _ := gScore[State{next, 0}]
			if next_g == 0 || (tentative_g < next_g) {
				gScore[State{next, 0}] = tentative_g
				heap.Push(openSet, Node{next, 0, tentative_g})
			}
			
		}
		
	}

	return 0
}

func display(mat [][]byte) {
	for _, row := range(mat) {
		for _, cell := range(row) {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}
