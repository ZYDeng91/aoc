package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	content, err := os.Open("1_converted.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	total := 0
	for scanner.Scan() {
		total += day1(scanner.Text())
	}
	fmt.Println(total)
}

func day1(text string) int {
	digit1 := -1
	digit2 := -1
	size := len(text)
	for i := 0; i < size; i++ {
		if digit1 == -1 && '0' <= text[i] && text[i] <= '9' {
			digit1 = int(text[i] - '0')
		}
		if digit2 == -1 && '0' <= text[size-i-1] && text[size-i-1] <= '9' {
			digit2 = int(text[size-i-1] - '0')
		}
		if digit1 != -1 && digit2 != -1 {
			return 10*digit1 + digit2
		}
	}
	return 0
}
