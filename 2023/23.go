package main

import (
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	x int
	y int
}
type Move struct {
	dist   int
	target Coord
}

func main() {
	content, err := os.ReadFile("23.in")
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")

	mat := make([][]rune, len(lines))

	for i, line := range lines {
		newline := make([]rune, len(line))
		for j, item := range line {
			newline[j] = item
		}
		mat[i] = newline
	}

	mat = mat[:len(mat)-1][:]

	fmt.Println(day23(mat))
}

func day23(mat [][]rune) int {
	start := Coord{0, 1}
	goal := Coord{len(mat) - 1, len(mat[0]) - 2}
	visited := make([][]int, len(mat))
	for i, line := range mat {
		newline := make([]int, len(line))
		for j, item := range line {
			if item == '#' {
				newline[j] = 2
				continue
			} else {
				newline[j] = 0
			}
		}
		visited[i] = newline
	}
	wormhole := make(map[Coord][]Move)

	res1 := step(0, start, mat, visited, Move{0, start}, wormhole)
	fmt.Println(res1)
	visited2 := make(map[Coord]bool)
	res2 := step2(0, start, goal, visited2, wormhole)
	fmt.Println(res2)
	return 0
}

func step2(i int, current, goal Coord, visited map[Coord]bool, wormhole map[Coord][]Move) int {
	visited[current] = true
	if current == goal {
		return i
	}
	next := make([]Move, 0)
	for _, newMove := range wormhole[current] {
		if visited[newMove.target] {
			continue
		}
		next = append(next, newMove)
	}

	if len(next) == 0 {
		return -1
	}

	record := 0
	for _, item := range next {
		record = max(record, step2(i+item.dist, item.target, goal, mapcpy(visited), wormhole))
	}
	return record
}

func step(i int, current Coord, mat [][]rune, visited [][]int, lastJunction Move, wormhole map[Coord][]Move) int {
	visited[current.x][current.y] = 1
	goal := Coord{len(mat) - 1, len(mat[0]) - 2}
	if current == goal {
		wormhole[lastJunction.target] = appendifnotin(wormhole[lastJunction.target], Move{i - lastJunction.dist, current})
		return i
	}
	dirs := []Coord{Coord{-1, 0}, Coord{0, 1}, Coord{1, 0}, Coord{0, -1}}
	next := make([]Coord, 0)
	for j, dir := range dirs {
		if (mat[current.x][current.y] == '^' && j != 0) || (mat[current.x][current.y] == '>' && j != 1) || (mat[current.x][current.y] == 'v' && j != 2) || (mat[current.x][current.y] == '<' && j != 3) {
			continue
		}
		newCoord := Coord{current.x + dir.x, current.y + dir.y}
		if newCoord.x < 0 || newCoord.x >= len(mat) || newCoord.y < 0 || newCoord.y >= len(mat[0]) {
			continue
		}
		if visited[newCoord.x][newCoord.y] > 0 {
			continue
		}

		next = append(next, newCoord)
	}

	if len(next) == 0 {
		return -1
	}

	if len(next) == 1 {
		return step(i+1, next[0], mat, arrcpy(visited), lastJunction, wormhole)
	}

	record := 0
	if len(next) > 1 {
		wormhole[current] = appendifnotin(wormhole[current], Move{i - lastJunction.dist, lastJunction.target})
		wormhole[lastJunction.target] = appendifnotin(wormhole[lastJunction.target], Move{i - lastJunction.dist, current})
	}
	for _, item := range next {
		record = max(record, step(i+1, item, mat, arrcpy(visited), Move{i, current}, wormhole))
	}

	return record
}

func arrcpy(visited [][]int) [][]int {
	visited_copy := make([][]int, len(visited))
	for i, temp1 := range visited {
		newline := make([]int, len(temp1))
		for j, temp2 := range temp1 {
			newline[j] = temp2
		}
		visited_copy[i] = newline
	}
	return visited_copy
}

func mapcpy(visited map[Coord]bool) map[Coord]bool {
	visited_copy := make(map[Coord]bool)
	for key, val := range visited {
		visited_copy[key] = val
	}
	return visited_copy
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func appendifnotin(b []Move, a Move) []Move {
	for _, item := range b {
		if a == item {
			return b
		}
	}
	return append(b, a)
}
