package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	content, err := os.Open("13.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	row := 0
	rows := make([][2]int, 3)

	res1, res2 := 0, 0

	for scanner.Scan() {
		text := scanner.Text()
		row += 1
		row = row % 4
		if row == 0 {
			res1 += day13(rows[0], rows[1], rows[2])
			res2 += day13p2(rows[0], rows[1], rows[2])
			continue
		}
		var nums_str []string
		if row == 3 {
			nums_str = strings.Split(text[9:], ", Y=")
		} else {
			nums_str = strings.Split(text[12:], ", Y+")
		}
		temp := make([]int, 2)
		for i, num_str := range nums_str {
			num, _ := strconv.Atoi(num_str)
			temp[i] = num
		}
		rows[row-1] = [2]int{temp[0], temp[1]}
	}
	fmt.Println(res1, res2)
}

func day13p2(a, b, prize [2]int) int {
	return day13(a, b, [2]int{prize[0]+10000000000000, prize[1]+10000000000000})
}

func day13(a, b, prize [2]int) int {
	nume := a[1]*prize[0] - a[0]*prize[1]
	denom := a[1]*b[0]-a[0]*b[1] 
	if denom == 0 {
		fmt.Println("division by zero", a, b, prize)
	}
	if nume%denom != 0 {
		return 0
	}
	x2 := nume/denom

	nume = prize[0]-x2*b[0]
	denom = a[0]

	if nume%denom != 0 {
		return 0
	}
	x1 := nume/denom

	if x1<0 || x2<0 {
		fmt.Println("underflow")
		fmt.Println(x1, x2, a, b, prize)
	}
	/*if x1>100 || x2>100 {
		fmt.Println("over limit")
		fmt.Println(x1, x2, a, b, prize)
	}*/
	return x1*3+x2*1
}
