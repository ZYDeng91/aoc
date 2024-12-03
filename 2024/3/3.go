package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	content, err := os.Open("3p2.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	total := 0
	total2 := 0
	do := 1
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}
		if text == "do()" {
			do = 1
			continue
		}
		if text == "don't()" {
			do = 0
			continue
		}
		text2 := text[4:len(text)-1]
		nums_str := strings.Split(text2, ",")
		nums := make([]int, 0)
		for _, num_str := range nums_str {
			num, _ := strconv.Atoi(num_str)
			nums = append(nums, num)
		}
		total += nums[0] * nums[1]
		total2 += do * nums[0] * nums[1]
	}
	fmt.Println(total, total2)
}
