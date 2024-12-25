package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type loc struct {
	x int
	y int
}

var numpad map[rune]loc
var dirpad map[rune]loc

type State struct {
	s rune
	e rune
	d int
}

var cached map[State]int

const maxInt int = int(^uint(0) >> 1)

func main() {
	content, err := os.Open("21.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)
	rows := make([]string, 0)

	for scanner.Scan() {
		text := scanner.Text()
		rows = append(rows, text)
	}

	res := 0

	numpad = make(map[rune]loc)
	dirpad = make(map[rune]loc)
	numpad['7'] = loc{0,0}
	numpad['8'] = loc{0,1}
	numpad['9'] = loc{0,2}
	numpad['4'] = loc{1,0}
	numpad['5'] = loc{1,1}
	numpad['6'] = loc{1,2}
	numpad['1'] = loc{2,0}
	numpad['2'] = loc{2,1}
	numpad['3'] = loc{2,2}
	numpad['0'] = loc{3,1}
	numpad['A'] = loc{3,2}
	dirpad['^'] = loc{0,1}
	dirpad['A'] = loc{0,2}
	dirpad['v'] = loc{1,1}
	dirpad['<'] = loc{1,0}
	dirpad['>'] = loc{1,2}

	cached = make(map[State]int)

	for _, row := range(rows) {
		num, _ := strconv.Atoi(row[:len(row)-1])
		//fmt.Println(day21(row))
		res += num*day21(row)
	}
	fmt.Println(res)
}

func day21(s string) int {
	dirs := num2dir(s)
	res := maxInt
	for _, d := range(dirs) {
		total := 0
		start := 'A'
		for _, end := range(d) {
			total += min_seq(start, end, 25)
			start = end
		}
		if res > total {
			res = total
		}
	}
	return res
}

func num2dir(s string) []string {
	start := 'A'
	res := []string{""}
	for _, end := range(s) {
		temp := make([]string, 0)
		nums := num(start, end)
		for _, item := range(res) {
			for _, num1 := range(nums) {
				temp = append(temp, item + num1)
			}
		}
		start = end
		res = temp
	}
	return res
}

func min_seq(start, end rune, depth int) int {
	val, ok := cached[State{start, end, depth}]
	if ok {
		return val
	}
	if depth == 0 {
		return 1
	}
	seqs := dir(start, end)
	/*if depth == 1 {
		return min_len(seqs)
	}*/

	res := maxInt

	for _, seq := range(seqs) {
		// note: a local var start/end does not conflict with the function inputs
		// probably an artifact of copypasting but it works
		start := 'A'
		temp := 0
		for _, end := range(seq) {
			temp += min_seq(start, end, depth-1)
			start = end
		}
		if res > temp {
			res = temp
		}
	}
	cached[State{start, end, depth}] = res
	return res
}

func min_len(a []string) int {
	res := maxInt
	for _, item := range(a) {
		if res > len(item) {
			res = len(item)
		}
	}
	return res
}

func num(start, end rune) []string {
	if start == end {
		return []string{"A"}
	}
	s := numpad[start]
	e := numpad[end]
	dx, dy := e.x - s.x, e.y - s.y
	if dx == 0 || dy == 0 {
		temp := make([]byte, abs(dx)+abs(dy)+1)
		var char byte
		if dx > 0 {
			char = 'v'
		} else {
			char = '^'
		}
		for i:=0;i<abs(dx);i++{
			temp[i] = char
		}
		if dy > 0 {
			char = '>'
		} else {
			char = '<'
		}
		for i:=0;i<abs(dy);i++{
			temp[abs(dx)+i] = char
		}
		temp[len(temp)-1] = 'A'
		return []string{string(temp)}
	} else {
		temp := make([]byte, abs(dx)+abs(dy)+1)
		temp2 := make([]byte, abs(dx)+abs(dy)+1)
		var char byte
		if dx > 0 {
			char = 'v'
		} else {
			char = '^'
		}
		for i:=0;i<abs(dx);i++{
			temp[i] = char
			temp2[abs(dy)+i] = char
		}
		if dy > 0 {
			char = '>'
		} else {
			char = '<'
		}
		for i:=0;i<abs(dy);i++{
			temp[abs(dx)+i] = char
			temp2[i] = char
		}
		temp[len(temp)-1] = 'A'
		temp2[len(temp2)-1] = 'A'
		if max(s.x, e.x) == 3 && min(s.y, e.y) == 0 {
			if char == '>' || char == '^' {
				return []string{string(temp2)}
			} else {
				return []string{string(temp)}
			}
		}
		return []string{string(temp), string(temp2)}
	}
}

func dir(start, end rune) []string {
	if start == end {
		return []string{"A"}
	}
	s := dirpad[start]
	e := dirpad[end]
	dx, dy := e.x - s.x, e.y - s.y
	if dx == 0 || dy == 0 {
		temp := make([]byte, abs(dx)+abs(dy)+1)
		var char byte
		if dx > 0 {
			char = 'v'
		} else {
			char = '^'
		}
		for i:=0;i<abs(dx);i++{
			temp[i] = char
		}
		if dy > 0 {
			char = '>'
		} else {
			char = '<'
		}
		for i:=0;i<abs(dy);i++{
			temp[abs(dx)+i] = char
		}
		temp[len(temp)-1] = 'A'
		return []string{string(temp)}
	} else {
		temp := make([]byte, abs(dx)+abs(dy)+1)
		temp2 := make([]byte, abs(dx)+abs(dy)+1)
		var char byte
		if dx > 0 {
			char = 'v'
		} else {
			char = '^'
		}
		for i:=0;i<abs(dx);i++{
			temp[i] = char
			temp2[abs(dy)+i] = char
		}
		if dy > 0 {
			char = '>'
		} else {
			char = '<'
		}
		for i:=0;i<abs(dy);i++{
			temp[abs(dx)+i] = char
			temp2[i] = char
		}
		temp[len(temp)-1] = 'A'
		temp2[len(temp2)-1] = 'A'
		if min(s.y, e.y) == 0 {
			if char == '>' || char == 'v' {
				return []string{string(temp2)}
			} else {
				return []string{string(temp)}
			}
		}
		return []string{string(temp), string(temp2)}
	}
}

func abs(a int) int {
	if a<0 {
		return -a
	}
	return a
}
