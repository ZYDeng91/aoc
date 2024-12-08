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
	content, err := os.ReadFile("8.in")
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")

	mat := make([]string, 0, len(lines))

	for _, line := range lines {
		mat = append(mat, line)
	}

	mat = mat[0 : len(mat)-1][:]

	fmt.Println(day8(mat))
}

func day8(mat []string) (int, int) {
	antennas := make(map[byte][]loc)
	for x := 0; x < len(mat); x++ {
		for y := 0; y < len(mat[x]); y++ {
			if mat[x][y] != '.' {
				antennas[mat[x][y]] = append(antennas[mat[x][y]], loc{x,y})
			}
		}
	}

	visited := makeEmpty(mat)
	max_x := len(mat)-1
	max_y := len(mat[0])-1

	for _, val := range(antennas) {
		for i:=0;i<len(val);i++ {
			for j:=i+1;j<len(val);j++ {
				anti1 := loc{2*val[i][0]-val[j][0], 2*val[i][1]-val[j][1]}
				anti2 := loc{2*val[j][0]-val[i][0], 2*val[j][1]-val[i][1]}
				if !(anti1[0]<0 || anti1[1]<0 || anti1[0]>max_x || anti1[1]>max_y){
					visited[anti1[0]][anti1[1]] = 1
				}

				if !(anti2[0]<0 || anti2[1]<0 || anti2[0]>max_x || anti2[1]>max_y){
					visited[anti2[0]][anti2[1]] = 1
				}
			}
		}
	}

	visited2 := makeEmpty(mat)
	for _, val := range(antennas) {
		for i:=0;i<len(val);i++ {
			for j:=i+1;j<len(val);j++ {
				diff := loc{val[i][0]-val[j][0], val[i][1]-val[j][1]}
				anti1 := val[i]
				anti2 := val[j]
				for !(anti1[0]<0 || anti1[1]<0 || anti1[0]>max_x || anti1[1]>max_y){
					visited2[anti1[0]][anti1[1]] = 1
					anti1[0] += diff[0]
					anti1[1] += diff[1]
				}

				for !(anti2[0]<0 || anti2[1]<0 || anti2[0]>max_x || anti2[1]>max_y){
					visited2[anti2[0]][anti2[1]] = 1
					anti2[0] -= diff[0]
					anti2[1] -= diff[1]
				}
			}
		}
	}

	return countVisited(visited),countVisited(visited2)
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
