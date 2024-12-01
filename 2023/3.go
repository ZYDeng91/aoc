package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("3.in")
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")

	mat := make([]string, 0, len(lines))

	for _, line := range lines {
		mat = append(mat, line)
	}

	fmt.Println(day3(mat))
}

func day3(mat []string) int {
	total := 0
	for x := 0; x < len(mat); x++ {
		for y := 0; y < len(mat[x]); y++ {
			if isNum(mat[x][y]) {
				top := x - 1
				if top < 0 {
					top = 0
				}
				bottom := x + 1
				if bottom == len(mat)-1 {
					bottom = len(mat) - 2
				}
				left := y - 1
				if left < 0 {
					left = 0
				}
				right := y + getNumLen(mat[x], y)
				if right == len(mat[x]) {
					right = len(mat[x]) - 1
				}
				if hasSymbol(mat, top, bottom, left, right) {
					total += getNum(mat[x], y)
					//fmt.Println(getNum(mat[x],y))
				}
				y += getNumLen(mat[x], y)
			}
		}
	}
	return total
}

func day3_2(mat []string) int {
	total := 0
	for x := 0; x < len(mat); x++ {
		for y := 0; y < len(mat[x]); y++ {
			if mat[x][y] == '*' {
				top := x - 1
				if top < 0 {
					top = 0
				}
				bottom := x + 1
				if bottom == len(mat)-1 {
					bottom = len(mat) - 2
				}
				left := y - 1
				if left < 0 {
					left = 0
				}
				right := y + 1
				if right == len(mat[x]) {
					right = len(mat[x]) - 1
				}
				total += hasNum(mat, top, bottom, left, right)
				y += getNumLen(mat[x], y)
			}
		}
	}
	return total
}

func isNum(a byte) bool {
	return '0' <= a && a <= '9'
}

func hasSymbol(mat []string, t, b, l, r int) bool {
	for x := t; x <= b; x++ {
		for y := l; y <= r; y++ {
			if (mat[x][y] != '.') && !isNum(mat[x][y]) {
				return true
			}
		}
	}
	return false
}

func hasNum(mat []string, t, b, l, r int) int {
	total := 1
	count := 0
	for x := t; x <= b; x++ {
		for y := l; y <= r; y++ {
			if isNum(mat[x][y]) {
				num := getNum(mat[x], getNumStart(mat[x], y))
				total *= num
				count += 1
				y = getNumStart(mat[x], y) + getNumLen(mat[x], y) + 1
			}
		}
	}
	if count == 2 {
		fmt.Println(total)
		fmt.Println(t, b, l, r)
		return total
	}
	return 0
}

func getNumLen(line string, loc int) int {
	if loc == len(line) {
		return 0
	}
	if isNum(line[loc]) {
		return getNumLen(line, loc+1) + 1
	}
	return 0
}

func getNum(line string, loc int) int {
	num, _ := strconv.Atoi(line[loc : loc+getNumLen(line, loc)])
	return num
}

func getNumStart(line string, loc int) int {
	if !isNum(line[loc]) {
		fmt.Println("Not a number, are you sure?")
		return -1
	}
	if loc == 0 || !isNum(line[loc-1]) {
		return loc
	}
	return getNumStart(line, loc-1)
}
