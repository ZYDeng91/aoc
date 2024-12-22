package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	content, err := os.Open("22.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)
	res := 0
	rows := make([]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		num, _ := strconv.Atoi(text)
		rows = append(rows, num)
		res += day22(num)
	}
	res2 := day22p2(rows)
	fmt.Println(res, res2)
}

func day22(n int) int {
	res := n
	for i:=0;i<2000;i++ {
		res = secret(res)
	}
	return res
}

func day22p2(ns []int) int {
	total := make(map[[4]int]int)
	for _, n := range(ns) {
		seen := make(map[[4]int]bool)
		seq := genseq(n)
		quad := [4]int{0,seq[1]-seq[0],seq[2]-seq[1],seq[3]-seq[2]}
		for i:=4;i<len(seq);i++ {
			quad = slide(quad, seq[i]-seq[i-1])
			if seen[quad] {
				continue
			} else {
				total[quad] += seq[i]
				seen[quad] = true
			}
		}
	}
	res := 0

	for _, v := range(total) {
		if v > res {
			res = v
		}
	}
	return res
}

func slide(quad [4]int, next int) [4]int {
	return [4]int{quad[1],quad[2],quad[3],next}
}

func genseq(n int) []int {
	res := make([]int, 2000)
	tmp := n
	for i:=0;i<len(res);i++ {
		res[i] = tmp%10
		tmp = secret(tmp)
	}
	return res
}

func secret(n int) int {
	return step3(step2(step1(n)))
}

func step1(n int) int {
	tmp := n * 64
	n = mix(n, tmp)
	n = prune(n)
	return n
}

func step2(n int) int {
	tmp := n / 32
	n = mix(n, tmp)
	n = prune(n)
	return n
}

func step3(n int) int {
	tmp := n * 2048
	n = mix(n, tmp)
	n = prune(n)
	return n
}

func mix(a, b int) int {
	return a^b
}

func prune(n int) int {
	return n%16777216
}
