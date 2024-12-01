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
	content, err := os.ReadFile("14.in")
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")

	mat := make([]string, 0, len(lines))

	//padding := strings.Repeat("#", len(lines[0]))
	//mat = append(mat, padding)

	for _, line := range lines {
		mat = append(mat, line)
	}

	mat = mat[0 : len(mat)-1][:]

	fmt.Println(day14_2(mat))
}

func day14(mat []string) int {
	piles := getPiles(mat, "N")
	total := 0
	for key, val := range piles {
		a := len(mat) - key.x - 1
		b := a - val + 1
		total += (a + b) * val / 2
	}
	return total
}

func getLoadN(mat [][]rune) int {
	total := 0
	for x := 0; x < len(mat); x++ {
		for y := 0; y < len(mat[x]); y++ {
			if mat[x][y] == 'O' {
				total += len(mat) - x
			}
		}
	}
	return total
}

func day14_2(mat []string) int {
	mat1 := str2rune2d(mat)
	// try first 1k rounds to look for pattern
	//num := 1000

	// repetition starts at round 180, each pattern is 14 rounds long
	num := (1000000000-180)%14 + 180
	for i := 1; i <= num; i++ {
		mat1 = round(mat1)
		fmt.Println(i, getLoadN(mat1))
	}
	return 0
}

func getStopsNS(mat []string) [][]int {
	stops := make([][]int, len(mat[0]))

	for x := 0; x < len(mat); x++ {
		for y := 0; y < len(mat[x]); y++ {
			if mat[x][y] == '#' {
				stops[y] = append(stops[y], x)
			}
		}
	}
	return stops
}
func getStopsWE(mat []string) [][]int {
	stops := make([][]int, len(mat[0]))

	for x := 0; x < len(mat); x++ {
		for y := 0; y < len(mat[x]); y++ {
			if mat[x][y] == '#' {
				stops[x] = append(stops[x], y)
			}
		}
	}
	return stops
}
func getPiles(mat []string, dir string) map[loc]int {
	stops := getStopsNS(mat)
	stops2 := getStopsWE(mat)
	piles := make(map[loc]int)
	for x := 0; x < len(mat); x++ {
		for y := 0; y < len(mat[x]); y++ {
			if mat[x][y] == 'O' {
				switch dir {
				case "N":
					target := -1
					for _, stop := range stops[y] {
						if x > stop {
							target = stop
						}
					}
					piles[loc{target, y}] += 1
					break
				case "S":
					target := len(mat)
					for i := len(stops[y]) - 1; i >= 0; i-- {
						if x < stops[y][i] {
							target = stops[y][i]
						}
					}
					piles[loc{target, y}] += 1
					break

				case "W":
					target := -1
					for _, stop := range stops2[x] {
						if y > stop {
							target = stop
						}
					}
					piles[loc{x, target}] += 1
					break
				case "E":
					target := len(mat[0])
					for i := len(stops2[x]) - 1; i >= 0; i-- {
						if y < stops2[x][i] {
							target = stops2[x][i]
						}
					}
					piles[loc{x, target}] += 1
					break
				}
			}
		}
	}
	return piles

}

func str2rune2d(mat []string) [][]rune {
	newArray := make([][]rune, len(mat))
	for x := 0; x < len(mat); x++ {
		newArray[x] = []rune(mat[x])
	}
	return newArray
}

func rune2str2d(mat [][]rune) []string {
	newArray := make([]string, len(mat))
	for x := 0; x < len(mat); x++ {
		newArray[x] = string(mat[x])
	}
	return newArray
}

func getSkeleton(mat [][]rune) [][]rune {
	after := make([][]rune, len(mat))
	for x := 0; x < len(mat); x++ {
		line := make([]rune, len(mat[0]))
		for y := 0; y < len(mat[0]); y++ {
			newSymbol := '.'
			if mat[x][y] == '#' {
				newSymbol = '#'
			}
			line[y] = newSymbol
		}
		after[x] = line
	}
	return after
}

func rotate(mat [][]rune, dir string) [][]rune {
	after := getSkeleton(mat)
	piles := getPiles(rune2str2d(mat), dir)
	for key, val := range piles {
		for i := 1; i <= val; i++ {
			if dir == "N" {
				after[key.x+i][key.y] = 'O'
			}
			if dir == "S" {
				after[key.x-i][key.y] = 'O'
			}
			if dir == "W" {
				after[key.x][key.y+i] = 'O'
			}
			if dir == "E" {
				after[key.x][key.y-i] = 'O'
			}
		}
	}
	return after
}

func round(mat [][]rune) [][]rune {
	mat = rotate(mat, "N")
	mat = rotate(mat, "W")
	mat = rotate(mat, "S")
	mat = rotate(mat, "E")
	return mat
}
