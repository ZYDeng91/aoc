package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	content, err := os.Open("23.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)
	res := 0
	rows := make([][2]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		left := text[:2]
		right := text[3:]
		rows = append(rows, [2]string{left, right})
	}
	res = day23(rows)
	fmt.Println(res)
}

func day23(pairs [][2]string) int {
	connected := make(map[string][]string)
	for _, pair := range(pairs) {
		l, r := pair[0], pair[1]
		connected[l] = append(connected[l], r)
		connected[r] = append(connected[r], l)
	}
	//fmt.Println(connected)
	res := 0

	for k, v := range(connected) {
		for i:=0;i<len(v)-1;i++ {
			for j:=i+1;j<len(v);j++ {
				if isin(v[j], connected[v[i]]) {
					if k[0] == 't' || v[i][0] == 't' || v[j][0] == 't' {
						res += 1
					}
					//fmt.Println(k,v[i],v[j])
				}
			}
		}
	}
	return res/3
}

func isin(item string, arr []string) bool {
	for _, i := range(arr) {
		if i == item {
			return true
		}
	}
	return false
}
