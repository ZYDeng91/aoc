package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	content, err := os.Open("24.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)
	mode := 0
	data := make(map[string]int)
	rows := make([][4]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			mode += 1
			continue
		}
		if mode == 0 {
			pair := strings.Split(text, ": ")
			num, _ := strconv.Atoi(pair[1])
			data[pair[0]] = num
		} else {
			inst := strings.Split(text, " ")
			rows = append(rows, [4]string{inst[0],inst[1],inst[2],inst[4]})
		}
	}
	//res := day24(data, rows)
	res2 := day24p2(data, rows)
	fmt.Println(res2)
}

func day24p2(data map[string]int, insts[][4]string) int {
	vocabs := make(map[string]string)
	temp := make([][4]string, 0)
	for {
		for _, inst := range(insts) {
			if (startsWith(inst[0],'x')||startsWith(inst[0],'y')) && (startsWith(inst[2],'y') || startsWith(inst[2],'x')){
				if inst[3][0] == 'z' {
					display(inst)
				} else {
					vocabs[inst[3]] = "("+inst[0]+" "+inst[1]+" "+inst[2]+")"
				}
			} else {
				temp = append(temp, inst)
			}
		}
		for i, inst := range(temp) {
			a, oka := vocabs[inst[0]]
			if oka {
				temp[i][0] = a
			}
			b, okb := vocabs[inst[2]]
			if okb {
				temp[i][2] = b
			}
		}
		if len(temp) == 0 {
			break
		} else {
			insts = temp
			temp = make([][4]string, 0)
		}
	}
	

	fmt.Println(temp)
	return 0
}

func startsWith(a string, b byte) bool {
	i := 0
	for {
		if a[i] == '(' {
			i+=1
		} else if a[i] == b {
			return true
		} else {
			return false
		}
	}
}

func display(row [4]string) {
	fmt.Println(row[0],row[1],row[2],"->",row[3])
}
