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
	content, err := os.Open("24p2.test")
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
	res := day24(data, rows)
	fmt.Println(res2)
}

func day24p2(data map[string]int, insts[][4]string) int {
	expected := output(data, 'x') + output(data, 'y')
	bin1 := dec2bin(expected)

	bin2 := dec2bin(day24(data, insts))
	for i, v := range(bin1) {
		if v != bin2[i] {
			fmt.Println(i)
		}
	}
	fmt.Println(bin1)
	fmt.Println(bin2)
	//expected := output(data, 'x') & output(data, 'y')
	//max_diff := bindiff(expected, day24(data, insts))
	/*for i:=0;i<len(insts)-1;i++{
		for j:=i+1;j<len(insts);j++{
			swapped := swap(insts, i, j)
			temp := day24(data, swapped)
			diff := bindiff(expected, temp)
			if diff < max_diff {
				fmt.Println(i, j, temp)
			}
		}
	}*/
	return 0
}

func cp(insts [][4]string) [][4]string {
	res := make([][4]string, len(insts))
	for i, inst := range(insts) {
		res[i] = [4]string{inst[0],inst[1],inst[2],inst[3]}
	}
	return res
}

func swap(insts [][4]string, i, j int) [][4]string {
	temp := cp(insts)
	temp[i][3], temp[j][3] = temp[j][3], temp[i][3]
	return temp
}

func bindiff(a, b int) int {
	dec := a ^ b
	bin := dec2bin(dec)
	res := 0
	for _, b := range(bin) {
		res += b
	}
	return res
}

// little endian
func dec2bin(dec int) []int {
	res := make([]int, 0)
	for dec > 0 {
		res = append(res, dec%2)
		dec = dec/2
	}
	return res
}

func day24(data2 map[string]int, insts [][4]string) int {
	data := make(map[string]int)
	for k, v := range(data2) {
		data[k] = v
	}
	for len(insts) > 0 {
		q := make([][4]string, 0)
		for _, inst := range(insts) {
			a, oka := data[inst[0]]
			b, okb := data[inst[2]]
			if !oka || !okb {
				q = append(q, inst)
				continue
			}
			data[inst[3]] = op(a, b, inst[1])
		}
		insts = q
	}
	return output(data, 'z')
}

func output(data map[string]int, wire byte) int {
	res := 0
	for k, v := range(data) {
		if k[0] == wire {
			i, _ := strconv.Atoi(k[1:])
			res += v*pow(2, i)
		}
	}
	return res
}

func pow(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func op(a, b int, gate string) int {
	switch gate {
		case "AND":
			return a & b
		case "OR":
			return a | b
		case "XOR":
			return a ^ b
		default:
			fmt.Println("error, ", gate)
			return 0
	}
}
