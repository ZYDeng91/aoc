package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    content, err := os.ReadFile("16.in")
    if err != nil {
    	return
    }

    lines := strings.Split(string(content), "\n")

    mat := make([]string, 0, len(lines))

    for _, line := range(lines){
   	mat = append(mat, line) 
    }

    mat = mat[0:len(mat)-1][:]

    fmt.Println(day16(mat))
}

func day16 (mat []string) (int, int) {
    var res []int
    visited := getEmptyMat(mat)
    for i:=0;i<len(mat);i++{
        visited = getEmptyMat(mat)
        start(i, -1, 'r', visited, mat, 0)
   	res = append(res, countVisited(visited))

        visited = getEmptyMat(mat)
        start(i, len(mat[0]), 'l', visited, mat, 0)
   	res = append(res, countVisited(visited))
    }

    for i:=0;i<len(mat[0]);i++{
        visited = getEmptyMat(mat)
        start(-1, i, 'd', visited, mat, 0)
   	res = append(res, countVisited(visited))

        visited = getEmptyMat(mat)
        start(len(mat), i, 'u', visited, mat, 0)
   	res = append(res, countVisited(visited))
    }

    visited = getEmptyMat(mat)
    start(0, -1, 'r', visited, mat, 0)

    return countVisited(visited), getMax(res)
}

func start (a, b int, dir rune, visited [][]int, mat []string, changed int) {
    // sometimes beams are overshadowed at first
    // and sometimes beams make loops
    for changed>-100{
	if dir == 'u'{
	    a -= 1
	}
	if dir == 'd'{
	    a += 1
	}
	if dir == 'l'{
	    b -= 1
	}
	if dir == 'r'{
	    b += 1
	}
	if a<0||b<0||a>=len(mat)||b>=len(mat[0]) {
	    break
	}
	if visited[a][b]!=1{
	    changed += 1
	    visited[a][b] = 1
	} else {
	    if changed > 0 {
		changed = 0
	    } else {
		changed -= 1
	    }
	}
	newdir, branch := getDir(dir, mat[a][b])
	if branch != '0' {
	    start(a, b, branch, visited, mat, changed)
	}
	dir = newdir
    }
}

func getDir(last rune, current byte) (rune, rune) {
    if last == 'u'{
	if current == '-' {
	    return 'l', 'r'
	}
	if current == '/' {
	    return 'r', '0'
	}
	if current == '\\' {
	    return 'l', '0'
	}
	return 'u', '0'
    }

    if last == 'd'{
	if current == '-' {
	    return 'l', 'r'
	}
	if current == '/' {
	    return 'l', '0'
	}
	if current == '\\' {
	    return 'r', '0'
	}
	return 'd', '0'
    }

    if last == 'l'{
	if current == '|' {
	    return 'u', 'd'
	}
	if current == '/' {
	    return 'd', '0'
	}
	if current == '\\' {
	    return 'u', '0'
	}
	return 'l', '0'
    }

    if last == 'r'{
	if current == '|' {
	    return 'u', 'd'
	}
	if current == '/' {
	    return 'u', '0'
	}
	if current == '\\' {
	    return 'd', '0'
	}
	return 'r', '0'
    }
    panic("invalid")
    return '0', '0'

}

func getEmptyMat (mat []string) [][]int {
    visited := make([][]int, 0, len(mat))
    for range(mat){
        emptyline := make([]int, 0, len(mat[0]))
        for range(mat[0]){
    	    emptyline = append(emptyline, 0)
        }
	visited = append(visited, emptyline)
    }
    return visited
}

func countVisited (visited [][]int) int {
    total := 0
    for _, row := range(visited) {
	for _, cell := range(row) {
	    if cell == 1 {
		total += 1
	    }
	}
    }
    return total
}

func getMax (arr []int) int {
    max := 0 
    for _, item := range(arr) {
	if item > max {
	    max = item
	}
    }
    return max
}
