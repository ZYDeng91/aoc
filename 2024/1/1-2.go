package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)
func main() {
	content, err := os.Open("1.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)
	left := make([]int, 0)
	right := make(map[int]int)

	re := regexp.MustCompile(`(\d+)\s+(\d+)`)

	for scanner.Scan() {
		text := scanner.Text()
		matches := re.FindStringSubmatch(text)
		left1, _  := strconv.Atoi(matches[1])
		right1, _  := strconv.Atoi(matches[2])
		left = append(left, left1)
		right[right1] += 1
	}
	fmt.Println(day1(left, right))
}

func day1(l []int, r map[int]int) int {
	total := 0
	i := len(l) - 1
	for i >= 0 {
		total += r[l[i]]*l[i]
		i--
	}
	return total
}
