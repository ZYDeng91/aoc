package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	//"strconv"
	//"strings"
	//"math"
)

func main() {
    content, err := os.Open("8.in")
    if err != nil {
    	return
    }
    defer content.Close()

    scanner := bufio.NewScanner(content)
    inst := ""
    lines := make(map[string]string)
    var ghosts []string

    for scanner.Scan(){
	text := scanner.Text()
	if text == "" {
	    continue
	}
	if inst == "" {
	    inst = text
	} else {
	    re := regexp.MustCompile(`(\w+) = \((\w+, \w+)\)`)
	    matches := re.FindStringSubmatch(text)
	    lines[matches[1]] = matches[2]
	    if matches[1][2] == 'A' {
		ghosts = append(ghosts, matches[1])
	    }
	}   
    }
    fmt.Println(day8(inst, lines))
    fmt.Println(day8_2(inst, lines, ghosts))
}

func day8 (inst string, lines map[string]string) int {
    sum := 0
    current := "AAA"
    for {
	for i, j := range(inst) {
	    current = getNext(j, lines[current])
	    if current == "ZZZ"{
		return sum+i+1
	    }
	}    
	sum += len(inst)
    }
    return -1
}


func day8_2 (inst string, lines map[string]string, ghosts []string) int {
    sum := 0

    // using direct copy will result in modifying the original
    current := make([]string, len(ghosts))
    for i, ghost := range(ghosts){
	current[i] = ghost
    }
    // brute force is taking too long
    // instead we observe periods when each ghost reach a Z 
    // and find least common multiple

    periods := make(map[string]int, len(ghosts))

    count := 0
    for a:=0;a<1000;a++{
	for i, j := range(inst) {
	    for p, q := range(current) {
	        current[p] = getNext(j, lines[q])
	        if current[p][2] == 'Z'{
		    count += 1
		    fmt.Printf("Ghost %s reached %s after %d steps \n", ghosts[p], current[p], sum+i+1)
		    // find a calculator or something to compute lcm
		    if periods[ghost[p]] == 0 {
			periods[ghost[p]] == sum+i+1
		    }
	        
	    	} 
	    }

	    if count == len(ghosts){
		return sum+i+1
	    }
	    count = 0
	}    
	sum += len(inst)
    }

    return -1
}

func getNext(inst rune, line string) string {
    if inst =='L'{
	return line[:3]
    }
    if inst =='R'{
	return line[5:]
    }
    return("cannot get next location")
} 
