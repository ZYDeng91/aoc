package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

var cached map[[2]int]int

func main() {
	content, err := os.Open("11.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	total := 0

	cached = make(map[[2]int]int)

	for scanner.Scan() {
		text := scanner.Text()
		nums_str := strings.Split(text, " ")
		//nums := make([]int, 0)
		for _, num_str := range nums_str {
			num, _ := strconv.Atoi(num_str)
			//nums = append(nums, num)
			total += day11(num, 75)
		}
	}
	fmt.Println(total)
}

func day11(num int, times int) int {
	if times == 0 {
		return 1
	}
	val, ok := cached[[2]int{num, times}]
	if ok {
		return val
	}

	res := 0

	after := blink(num)
	for _, item := range(after) {
		res += day11(item, times-1)
	}

	cached[[2]int{num, times}] = res
	return res
}

func getDigits(a int) int {
	b := 1
	for pow(10, b) <= a {
		b += 1
	}
	return b
}

func pow(a, b int) int {
	res := a
	for i:=1;i<b;i++{
		res *= a
	}
	return res
}

func split(ab int) (int, int) {
	half := pow(10, getDigits(ab)/2)
	a := ab / half
	b := ab % half
	return a, b
}

func blink(a int) []int {
	if a == 0 {
		return []int{1}
	}
	if getDigits(a)%2 == 0 {
		b, c := split(a)
		return []int{b, c}
	}

	return []int{a*2024}
}
