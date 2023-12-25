package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
    content, err := os.Open("25.in")
    if err != nil {
    	return
    }
    defer content.Close()

    scanner := bufio.NewScanner(content)

    lines := make(map[string][]string)
    comp := make([]string, 0)

    for scanner.Scan(){
	text := scanner.Text()
	splits := strings.Split(text, ": ")
	lines[splits[0]] = strings.Split(splits[1], " ")
	comp = append(comp, splits[0])
    }
    for key, val := range(lines) {
	for _, item := range(val) {
	    if len(lines[item]) == 0 {
		comp = append(comp, item)
	    }
	    lines[item] = append(lines[item], key)
	}
    }
    fmt.Println(day25(comp, lines))
}

func day25 (components []string, connections map[string][]string) int {
    group := make(map[string]bool)
    start := components[0]
    group[start] = true
    vs := []string{start}
    for len(vs)>0{
	newV:=make([]string, 0)
	for _, v := range(vs) {
	    for _, conn := range(connections[v]) {
		if (v == "zhb" && conn == "vxr") || (v == "vxr" && conn == "zhb") {
		    continue
		}
		if (v == "jbx" && conn == "sml") || (v == "sml" && conn == "jbx") {
		    continue
		}
		if (v == "szh" && conn == "vqj") || (v == "vqj" && conn == "szh") {
		    continue
		}
		if group[conn] {
		    continue
		}
		group[conn] = true
		newV = append(newV, conn)
	    }
	}
	vs = newV
    }
    return len(group)*(len(components)-len(group))
}
