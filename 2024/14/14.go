package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Bot struct {
	p [2]int
	v [2]int
}

func main() {
	content, err := os.Open("14.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	total := 0

	p := make([][]int, 0)
	v := make([][]int ,0)
	for scanner.Scan() {
		text := scanner.Text()
		nums_str := strings.Split(text[2:], " v=")

		left := make([]int, 2)
		right := make([]int, 2)

		for i, num_str := range nums_str {
			num_str_split := strings.Split(num_str, ",")
			for j, k := range(num_str_split) {
				val, _ := strconv.Atoi(k)
				if i == 0 {
					left[j] = val
				} else {
					right[j] = val
				}
			}
		}
		p = append(p, left)
		v = append(v, right)
	}

	total = day14(p, v, 100, 101, 103)//11, 7)
	day14p2(p, v, 20000, 101, 103)
	fmt.Println(total)
}

func day14(p, v [][]int, t int, a, b int) int {
	tl, tr, br, bl := 0,0,0,0
	halfa := (a-1)/2
	halfb := (b-1)/2
	for i:=0;i<len(p);i++ {
		final := [2]int{p[i][0]+v[i][0]*t, p[i][1]+v[i][1]*t}
		finala := (final[0]%a+a)%a - halfa
		finalb := (final[1]%b+b)%b - halfb
		if finala < 0 {
			if finalb < 0 {
				tl += 1
			} else if finalb > 0 {
				bl += 1
			}
		} else if finala > 0 {
			if finalb < 0 {
				tr += 1
			} else if finalb > 0 {
				br += 1
			}
		}
	}
	return tl*tr*br*bl
}

func makeEmpty(a, b int) [][]int {
	empty := make([][]int, b)
	for i:=0;i<b;i++{
		row := make([]int, a)
		empty[i] = row
	}
	return empty
}

func day14p2(p, v [][]int, time, a, b int) int {
	bots := make([]Bot, len(p))
	for i:=0;i<len(p);i++ {
		// use row, col instead of a, b
		bots[i] = Bot{[2]int{p[i][1],p[i][0]}, [2]int{v[i][1],v[i][0]}}
	}

	state := makeEmpty(a, b)
	for _, bot := range(bots) {
		state[bot.p[0]][bot.p[1]] = 1
	}

	for t:=0;t<time;t++ {
		fmt.Println("time:", t)
		display(state)
		//refresh(b)

		for i, bot := range(bots) {
			state[bot.p[0]][bot.p[1]] -= 1
			bot.p[0] = (bot.p[0]+bot.v[0]+b)%b
			bot.p[1] = (bot.p[1]+bot.v[1]+a)%a
			state[bot.p[0]][bot.p[1]] += 1
			bots[i] = bot
		}
	}
	return 0
}

func display(mat [][]int) {
	for _, row := range(mat) {
		for _, cell := range(row) {
			if cell == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Println()
	}
}

func refresh(rows int) {
	for i:=0;i<rows;i++ {
		fmt.Print("\033[1A\033[2K\r")
	}
}
