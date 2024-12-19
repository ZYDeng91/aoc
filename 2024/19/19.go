package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cached map[string]bool
var cached2 map[string]int

func main() {
	content, err := os.Open("19.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	section := 0
	patterns := make([]string, 0)
	designs := make([]string, 0)

	res1, res2 := 0, 0

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			section += 1
			continue
		}
		if section == 0 {
			patterns = strings.Split(text, ", ")
		}
		if section == 1 {
			designs = append(designs, text)
		}
	}
	patterns2 := make(map[string]bool)
	max_pattern := 0
	for _, pattern := range(patterns) {
		patterns2[pattern] = true
		if len(pattern) > max_pattern {
			max_pattern = len(pattern)
		}
	}
	cached = make(map[string]bool)
	cached2 = make(map[string]int)
	res1, res2 = day19(patterns2, designs, max_pattern)
	fmt.Println(res1, res2)
}

func day19(p map[string]bool, ds []string, m int) (int, int) {
	res := 0
	res2 := 0
	for _, d := range(ds) {
		if next(p, d, m) {
			res += 1
		}
		res2 += next2(p, d, m)
	}
	return res, res2
}

func next(p map[string]bool, d string, m int) bool {
	val, ok := cached[d]
	if ok {
		return val
	}
	if len(d) <= m && p[d] {
		cached[d] = true
		return true
	}
	for n:=1;n<=min(m,len(d));n++ {
		if p[d[:n]] {
			if next(p, d[n:], m) {
				cached[d] = true
				return true
				break
			}
		}
	}
	cached[d] = false
	return false
}

func next2(p map[string]bool, d string, m int) int {
	val, ok := cached2[d]
	if ok {
		return val
	}
	res := 0
	if p[d] {
		res += 1
	}
	for n:=1;n<=min(m,len(d));n++ {
		if p[d[:n]] {
			res += next2(p, d[n:], m)
		}
	}
	cached2[d] = res
	return res
}
