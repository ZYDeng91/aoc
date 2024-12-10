package main

import (
	"fmt"
	"os"
	"strings"
)

type loc [2]int

type state struct {
	pos loc
	dir loc
}

var visited map[loc]bool

func main() {
	content, err := os.ReadFile("10.in")
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")

	mat := make([]string, 0, len(lines))

	for _, line := range lines {
		mat = append(mat, line)
	}

	mat = mat[0 : len(mat)-1][:]

	fmt.Println(day10(mat))
}

func day10(mat []string) (int, int) {
	res := 0
	res2 := 0
	heads := make([]loc, 0)
	max_x := len(mat)
	max_y := len(mat[0])
	for x := 0; x < max_x; x++ {
		for y := 0; y < max_y; y++ {
			if mat[x][y] == '0' {
				heads = append(heads, loc{x, y})
			}
		}
	}

	for _, head := range(heads) {
		visited = make(map[loc]bool)
		res2 += path(mat, head, '0', max_x, max_y)
		temp := countVisited(visited)
		//fmt.Println(temp, head)
		res += temp
	}

	return res, res2
}

func path(mat []string, current loc, state byte, max_x, max_y int) int {
	if state == '9' {
		visited[current] = true
		return 1
	}
	res := 0
	dirs := make([]loc, 4)
	for i:=0;i<4;i++ {
		dirs[i] = loc{current[0]+i%2*(2*(i/2)-1), current[1]+(i-1)%2}
	}
	for _, dir := range(dirs) {
		if !(dir[0] < 0 || dir[0] >= max_x || dir[1] < 0 || dir[1] >= max_y) {
			if mat[dir[0]][dir[1]] == state + 1 {
				res += path(mat, dir, state+1, max_x, max_y)
			}
		}
	}
	return res
}

func countVisited(visited map[loc]bool) int{
	res := 0
	for _, v := range(visited) {
		if v {
			res += 1
		}
	}
	return res
}
/*
func makeEmpty(mat []string) [][]int {
	visited := make([][]int, 0, len(mat))
	for range mat {
		emptyline := make([]int, 0, len(mat[0]))
		for range mat[0] {
			emptyline = append(emptyline, 0)
		}
		visited = append(visited, emptyline)
	}
	return visited
}*/
/*
func cp(mat []string) []string {
	temp := make([]string, len(mat))
	for x := 0; x < len(mat); x++ {
		line := make([]byte, len(mat))
		for y := 0; y < len(mat[0]); y++ {
			line[y] = mat[x][y]
		}
		temp[x] = line
	}
	return temp
}*/
