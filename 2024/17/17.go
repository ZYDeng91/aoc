package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"strconv"
)

func main() {
	content, err := os.Open("17.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	row := -1
	rows := make([]int, 3)

	program := make([]int, 0)

	//res1, res2 := 0, 0

	for scanner.Scan() {
		row += 1
		text := scanner.Text()
		if row == 3 {
			continue
		}
		if row == 4 {
			nums_str := strings.Split(text[9:], ",")
			for _, num_str := range nums_str {
				num, _ := strconv.Atoi(num_str)
				program = append(program, num)
			}
			continue
		}
		num_str := text[12:]
		num, _ := strconv.Atoi(num_str)
		rows[row] = num
	}
	day17(rows, program)

	day17p2(program)
	//fmt.Println(res1, res2)
}

func day17(rs, p []int) int {
	a := rs[0]
	b := rs[1]
	c := rs[2]
	for i:=0;i<len(p);i+=2 {
		mainloop:
		
		opcode := p[i]
		operand_sel := p[i+1]
		operand := 0
		switch operand_sel {
			case 4:
				operand = a
			case 5:
				operand = b
			case 6:
				operand = c
			case 7:
				//fmt.Println("error", i)
			default:
				operand = operand_sel
		}

		switch opcode {
			case 0:
				a = a / pow(2,operand)
			case 1:
				b = b ^ operand_sel
			case 2:
				b = operand % 8
			case 3:
				if a != 0 {
					i = operand_sel
					//fmt.Println("jumping to", i)
					goto mainloop
				}
			case 4:
				b = b ^ c
			case 5:
				fmt.Print(operand%8, ",")
			case 6:
				b = a / pow(2,operand)
			case 7:
				c = a / pow(2,operand)
		}
	}
	fmt.Println()
	return 0
}

// little endian
func day17p2(p []int) {
	a := make([]int, len(p)*3+7)
	ptr := 7

	res := 0

	for i:=0;i<len(p);i++{
		ptr += 3
		for j:=0;j<8;j++{
			a[ptr-3], a[ptr-2], a[ptr-1] = oct2bin(j)
			//fmt.Println(ptr-(j^2))
			//fmt.Println(j, bin2oct([3]int(a[ptr-3-(j^2):ptr-(j^2)])), j^5^bin2oct([3]int(a[ptr-(j^2)-3:ptr-(j^2)])))
			if j^5^bin2oct([3]int(a[ptr-(j^2)-3:ptr-(j^2)])) == p[len(p)-i-1] {
				res = res * 8
				res += j
				//fmt.Print(j)
				break
			}
		}
	}
	fmt.Println(res)
}

func bin2oct(a [3]int) int {
	return 4*a[0]+2*a[1]+a[2]
}

func oct2bin(a int) (int, int, int) {
	return a/4, a/2-a/4*2, a%2
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
