package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	content, err := os.Open("23.ex")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)
	rows := make([][2]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		left := text[:2]
		right := text[3:]
		rows = append(rows, [2]string{left, right})
	}
	day23p2(rows)
}
func day23p2(pairs [][2]string) {
	fmt.Print("G.add_edges_from([")
	for i, pair := range(pairs) {
		l, r := pair[0], pair[1]
		if i != 0 {
			fmt.Print(",(\""+l+"\",\""+r+"\")")
		} else {
			fmt.Print("(\""+l+"\",\""+r+"\")")
		}
	}
	fmt.Println("])")
	/*connected := make(map[string][]string)
	for _, pair := range(pairs) {
		l, r := pair[0], pair[1]
		connected[l] = append(connected[l], r)
	}
	fmt.Println("graph {")
	for k, v := range(connected) {
		fmt.Print(k)
		fmt.Print(" -- {")
		for i, n := range(v) {
			if i != 0 {
				fmt.Print(" ")
			}
			fmt.Print(n)
		}
		fmt.Println("}")
	}
	fmt.Println("}")*/
}
