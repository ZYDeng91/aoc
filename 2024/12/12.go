package main

import (
	"fmt"
	"os"
	"strings"
)

type loc [2]int

type Fence [2][2]int

var seen map[loc]bool
var fences map[Fence]bool

func main() {
	content, err := os.ReadFile("12.in")
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")

	mat := make([]string, 0, len(lines))

	for _, line := range lines {
		mat = append(mat, line)
	}

	mat = mat[0 : len(mat)-1][:]

	fmt.Println(day12(mat))
}

func day12(mat []string) (int, int) {
	res := 0
	res2 := 0
	max_x := len(mat)
	max_y := len(mat[0])

	visited := makeEmpty(mat)
	
	for x := 0; x < max_x; x++ {
		for y := 0; y < max_y; y++ {
			if visited[x][y] == 0 {
				seen = make(map[loc]bool)
				fences = make(map[Fence]bool)
				temp1, temp2, temp3 := getRegion(mat, [2]int{x, y}, mat[x][y], max_x, max_y)
				res += temp1 * temp2
				res2 += temp1 * temp3
				for key, val := range(seen) {
					if val {
						visited[key[0]][key[1]] = 1
					}
				}
			}
		}
	}

	return res, res2
}

func getRegion(mat []string, current loc, state byte, max_x, max_y int) (int, int, int) {
	seen[current] = true

	res1, res2, res3 := 1, 0, 0
	next := make([]loc, 0)
	dirs := make([]loc, 4)
	for i:=0;i<4;i++ {
		dirs[i] = loc{current[0]+(i-2)%2, current[1]+(i-1)%2}
	}
	for _, dir := range(dirs) {
		if !(dir[0] < 0 || dir[0] >= max_x || dir[1] < 0 || dir[1] >= max_y) {
			if seen[dir] {
				continue
			} else if mat[dir[0]][dir[1]] == state {
				next = append(next, dir)
				continue
			}
		}
		res2 += 1
		fence_dir := loc{dir[0]-current[0],dir[1]-current[1]}
		fence := Fence{current, fence_dir}
		fences[fence] = true
		fence1 := Fence{loc{current[0]-fence_dir[1], current[1]-fence_dir[0]}, fence_dir}

		fence2 := Fence{loc{current[0]+fence_dir[1], current[1]+fence_dir[0]}, fence_dir}

		// this side is counted twice
		if fences[fence1] && fences[fence2] {
			res3 -= 1
			continue
		}
		if fences[fence1] || fences[fence2] {
			continue
		}

		res3 += 1
	}
	for _, i := range(next) {
		if seen[i] {
			continue
		}
		temp1, temp2, temp3 := getRegion(mat, i, state, max_x, max_y)
		res1 += temp1
		res2 += temp2
		res3 += temp3
	}

	//fmt.Println(state, res1, res2, res3)
	
	
	return res1, res2, res3
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
}
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
