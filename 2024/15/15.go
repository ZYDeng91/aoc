package main

import (
	"fmt"
	"os"
	"strings"
)

type loc struct {
	x int
	y int
}

func main() {
	content, err := os.ReadFile("15.in")
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")

	mat := make([][]byte, 0)
	moves := make([]byte, 0)

	mode := 0

	for _, line := range lines {
		if len(line) == 0 {
			mode = 1
		}
		if mode == 0 {
			mat = append(mat, []byte(line))
		} else {
			moves = append(moves, line...)
		}
	}

	fmt.Println(day15(mat, moves))
}

func day15(mat [][]byte, moves []byte) int {
	res := 0
	var start loc
	for x, row := range(mat) {
		for y, cell := range(row) {
			if cell == '@' {
				start = loc{x, y}
			}
		}
	}

	for _, move := range(moves) {
		dir := loc{0, 0}
		switch move {
			case '^':
				dir = loc{-1,0}
			case '>':
				dir = loc{0,1}
			case 'v':
				dir = loc{1,0}
			case '<':
				dir = loc{0,-1}
			default:
				fmt.Println("Invalid move")
		}
		target := mat[start.x+dir.x][start.y+dir.y]
		if target == '#' {
			continue
		} else if target == 'O' {
			for i:=2;i<100;i++{
				target2 := mat[start.x+dir.x*i][start.y+dir.y*i]
				if target2 == 'O' {
					continue
				} else if target2 == '#' {
					break
				} else {
					mat[start.x+dir.x*i][start.y+dir.y*i] = 'O'
					start = loc{start.x+dir.x, start.y+dir.y}
					mat[start.x][start.y] = '.'
					break
				}
			}
		} else {
			start = loc{start.x+dir.x, start.y+dir.y}
		}
	}

	for x, row := range(mat) {
		for y, cell := range(row) {
			if cell == 'O' {
				res += 100*x+y
			}
		}	
	}

	return res
}
