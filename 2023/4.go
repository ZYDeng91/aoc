package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	//"strconv"
	"math"
	"strings"
)

func main() {
	content, err := os.Open("4.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	total := 0
	index := 0
	multi := make(map[int]int)

	for scanner.Scan() {
		index += 1
		multi[index] += 1
		wins := day4(scanner.Text())
		for i := 1; i <= wins; i++ {
			multi[index+i] += multi[index]
		}
		total += multi[index]
	}
	fmt.Println(total)
}

func day4(text string) int {
	split := regexp.MustCompile(`.*: (.*) \| (.*)`)
	rm := regexp.MustCompile(`\s+`)
	line := split.FindStringSubmatch(rm.ReplaceAllString(text, " "))
	left := strings.Split(line[1], " ")
	right := strings.Split(line[2], " ")
	matches := 0
	for _, l := range left {
		for _, r := range right {
			if l == r {
				matches += 1
			}
		}
	}
	return matches
	if matches != 0 {
		return int(math.Pow(2, float64(matches-1)))
	}

	return 0
}
