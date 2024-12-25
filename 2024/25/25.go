package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	content, err := os.Open("25.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)
	mode := 0
	row := 0
	locks := make([][5]int, 0)
	pins := make([][5]int, 0)
	var item [5]int
	for scanner.Scan() {
		text := scanner.Text()
		row += 1
		if text == "" {
			continue
		}
		if row%8 == 1 {
			if text == "#####" {
				mode = 0
				item = [5]int{0,0,0,0.0}
			}
			if text == "....." {
				mode = 1
				item = [5]int{-1,-1,-1,-1,-1}
			}
			continue
		}
		for i, b := range(text) {
			if mode == 0 && b == '#' {
				item[i] += 1
			}
			if mode == 1 && b == '#' {
				item[i] += 1
			}
		}

		if row%8 == 7 {
			if mode == 0 {
				locks = append(locks, item)
			} else {
				pins = append(pins, item)
			}
		}
	}
	res := day25(locks, pins)
	fmt.Println(res)
}

func day25(locks, pins [][5]int) int {
	res := 0
	for _, lock := range(locks) {
		for _, pin := range(pins) {
			if check(lock, pin) {
				res += 1
			}
		}
	}
	return res
}

func check(lock, pin [5]int) bool {
	for i:=0;i<5;i++{
		if lock[i]+pin[i]>5 {
			return false
		}
	}
	return true
}
