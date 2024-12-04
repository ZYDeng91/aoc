package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("4.in")
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")

	mat := make([]string, 0, len(lines))

	for _, line := range lines {
		mat = append(mat, line)
	}

	mat = mat[0 : len(mat)-1][:]

	fmt.Println(day4(mat, "XMAS"), day4p2(mat))
}

func day4(mat []string, word string) int {
	res := 0
	for x := 0; x < len(mat); x++ {
		for y := 0; y < len(mat[0]); y++ {
			if mat[x][y] == word[0] {
				dirs := getDirs(mat, len(word)-1, x, y)
				for _, val := range(dirs) {
					if val == word {
						res += 1
					}
				}
			}
		}
	}
	return res
}

func getDirs(mat []string, l int, x, y int) []string {
	res := make([]string, 0)
	max_x := len(mat)-1
	max_y := len(mat[0])-1
	for _, xx := range []int{-1,0,1} {
		for _, yy := range[]int{-1,0,1} {
			if xx == 0 && yy == 0 {
				continue
			}
			if x+xx*l <= max_x && x+xx*l >= 0 && y+yy*l <= max_y && y+yy*l >= 0 {
				dir := make([]byte, l+1)
				for n:=0;n<=l;n++ {
					dir[n] = mat[x+xx*n][y+yy*n]
				}
				res = append(res, string(dir))
			} 
		}
	}
	return res
}

func day4p2(mat []string) int {
	res := 0
	for x := 0; x < len(mat); x++ {
		for y := 0; y < len(mat[0]); y++ {
			if mat[x][y] == 'A' {
				if x != 0 && y != 0 && x != len(mat)-1 && y != len(mat[0])-1 {
					res += check(mat, x, y)
				}
			}
		}
	}
	return res

}

func check(mat []string, x, y int) int {
	tl := rune(mat[x-1][y-1])
	tr := rune(mat[x-1][y+1])
	bl := rune(mat[x+1][y-1])
	br := rune(mat[x+1][y+1])

	MMSS := int('M'*'M'*'S'*'S') 
	if int(tl * tr * bl * br) != MMSS {
		return 0
	}

	if tl * br == 'M'*'M' || tr * bl == 'M'*'M' {
		return 0
	}
	return 1
}
