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

var pushed []loc

func main() {
	content, err := os.ReadFile("15p2.in")
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
				mat[x][y] = '.'
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
		//display(mat, start)
		if target == '#' {
			continue
		} else if target == '.' {
			start = loc{start.x+dir.x, start.y+dir.y}
			continue
		}
		targets := getBox(mat, loc{start.x+dir.x, start.y+dir.y})
		pushed = make([]loc, 0)	
		if pushable(mat, targets, dir) {
			//fmt.Println(pushed)
			mat = push(mat, dir)
			start = loc{start.x+dir.x, start.y+dir.y}
		}

	}

	for x, row := range(mat) {
		for y, cell := range(row) {
			if cell == '[' {
				res += 100*x+y
			}
		}	
	}

	return res
}

func clone(mat [][]byte) [][]byte {
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

func pushable(mat [][]byte, target []loc, dir loc) bool {
	pushed = append(pushed, target...)
	res := true
	for _, t := range(target) {
		temp := mat[t.x+dir.x][t.y+dir.y]
		if temp == '#' {
			return false
		} else if temp == '.' {
			continue
		} else {
			skip := false
			for _, t2 := range(target) {
				if t2.x == t.x + dir.x && t2.y == t.y + dir.y {
					skip = true
				}
			}
			//fmt.Println(getBox(mat, loc{t.x+dir.x,t.y+dir.y}), string(mat[t.x+dir.x][t.y+dir.y]))
			if !skip {
				res = res && pushable(mat, getBox(mat, loc{t.x+dir.x,t.y+dir.y}), dir)
			}
		}
	}
	return res
}

func getBox(mat [][]byte, pos loc) []loc {
	if mat[pos.x][pos.y] == '[' {
		return []loc{pos, loc{pos.x, pos.y+1}}
	} else {
		return []loc{pos, loc{pos.x, pos.y-1}}
	}
}

func push(mat [][]byte, dir loc) [][]byte {
	//fmt.Println(dir, pushed)
	res := clone(mat)
	for _, item := range(pushed) {
		res[item.x][item.y] = '.'
	}
	for _, item := range(pushed) {
		res[item.x+dir.x][item.y+dir.y] = mat[item.x][item.y]
	}
	return res
}

func display(a [][]byte, s loc) {
	for x, row := range(a) {
		for y, cell := range(row) {
			if x == s.x && y == s.y {
				fmt.Print("@")
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}
}
