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

func main() {
	content, err := os.ReadFile("6.in")
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")

	mat := make([]string, 0, len(lines))

	for _, line := range lines {
		mat = append(mat, line)
	}

	mat = mat[0 : len(mat)-1][:]

	fmt.Println(day6(mat))
}

func day6(mat []string) (int, int) {

	dir := loc{-1,0}
	last := loc{0,0}
	start := loc{0,0}
	for x := 0; x < len(mat); x++ {
		for y := 0; y < len(mat[x]); y++ {
			if mat[x][y] == '^' {
				start = loc{x,y}
				last = loc{x,y}
				break
			}
		}
	}

	visited := makeEmpty(mat)

	for k := 0; k < 100000; k++ {
		visited[last[0]][last[1]] = 1
		next := loc{last[0]+dir[0], last[1]+dir[1]}
		if next[0] < 0 || next[0] >= len(mat) || next[1] < 0 || next[1] >= len(mat[0]) {
			break
		}
		if mat[next[0]][next[1]] == '#' {
			dir = getDir(dir)
			continue
		}
		last = next
	}

	//return countVisited(visited)
	p1 := countVisited(visited)

	p2 := 0
	for x := 0; x < len(visited); x++ {
		for y := 0; y < len(visited[0]); y++ {
			if visited[x][y] == 1 {
				if x == start[0] && y == start[1] {
					continue	
				}
				obstacle := loc{x,y}
				visited2 := makeEmpty(mat)
				last = start
				dir = loc{-1, 0}
				for k := 0; k < 100000; k++ {
					visited2[last[0]][last[1]] += 1
					if visited2[last[0]][last[1]] > 4 {
					//if k == 99999 {
						p2 += 1
						break
					}
					next := loc{last[0]+dir[0], last[1]+dir[1]}
					if next[0] < 0 || next[0] >= len(mat) || next[1] < 0 || next[1] >= len(mat[0]) {
						break
					}
					if mat[next[0]][next[1]] == '#' || next == obstacle {
						dir = getDir(dir)
						continue
					}
					last = next
				}

			}
		}
	}
	return p1, p2
	//visited2 := make(map[state]bool)
}

func getDir(last loc) loc {
	switch last {
	case loc{-1, 0}:
		return loc{0, 1}
	case loc{0, 1}:
		return loc{1, 0}
	case loc{1, 0}:
		return loc{0, -1}
	case loc{0, -1}:
		return loc{-1, 0}
	default:
		fmt.Println("invalid dir")
		return loc{-1, -1}
	}
}

func countVisited (mat [][]int) int {
	res := 0
	for x := 0; x < len(mat); x++ {
		for y := 0; y < len(mat[0]); y++ {
			if mat[x][y] == 1 {
				res += 1
			}
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
