package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

//type Workflow func(Part)string
type Workflow []string

type Part map[byte]int

func main() {
    content, err := os.Open("19.in")
    if err != nil {
    	return
    }
    defer content.Close()

    scanner := bufio.NewScanner(content)

    workflows := make(map[string]Workflow)
    lines := make([]Part, 0)

    for scanner.Scan(){
	text := scanner.Text()
	if len(text) == 0{
	    break
	}
	i := 0
	for j, char := range(text) {
	    if char == '{' {
		i = j
		break
	    }
	}
	name := text[:i]
	splits := strings.Split(text[i+1:len(text)-1], ",")
    	workflows[name] = splits
    }
    for scanner.Scan(){
	text := scanner.Text()
	splits := strings.Split(text[1:len(text)-1], ",")
	line := make(Part)
	for _, item := range(splits) {
	    num, _ := strconv.Atoi(item[2:])
            line[item[0]] = num
        }
	lines = append(lines, line)
    }

    fmt.Println(day19(workflows, lines))
    fmt.Println(day19_2(workflows))
}

func day19_2 (workflows map[string]Workflow) int {
    start := make(Part)
    end := make(Part)
    for _, char := range([]byte{'x','m','a','s'}) {
	start[char] = 1
	end[char] = 4000
    }
    res := "in"
    return branch(start, end, workflows, res, 0)
}

func branch (start, end Part, workflows map[string]Workflow, res string, i int) int{
    if res == "A" {
	return (end['x']-start['x']+1)*(end['m']-start['m']+1)*(end['a']-start['a']+1)*(end['s']-start['s']+1)
    }
    if res == "R" {
	return 0
    }
    if !isValid(start, end){
	return 0
    }
    if i == len(workflows[res])-1 {
        start1, end1, res1 := filter(start, end, workflows[res][i], true)
        return branch(start1, end1, workflows, res1, 0) 
    }

    start1, end1, res1 := filter(start, end, workflows[res][i], true)
    start2, end2, _ := filter(start, end, workflows[res][i], false)

    return branch(start1, end1, workflows, res1, 0) + branch(start2, end2, workflows, res, i+1)
}


func day19 (workflows map[string]Workflow, parts []Part) int{
    total := 0 
    for _, part := range(parts) {
	workflow := workflows["in"]
        //fmt.Println(getWorkflow(workflow, part))
	for a:=0;a<1000;a++{
	    res := getWorkflow(workflow, part)
	    if res == "A" {
                total += part['x']+part['m']+part['a']+part['s']
		break
	    }
	    if res == "R" {
	        break
	    }
	    workflow = workflows[res]
        }
    
    }
    return total
}

func getOp (op string, part Part) string {
    split := strings.Split(op, ":")
    if len(split) == 1 {
	return op
    }
    cond := split[0]
    res := split[1]
    if check(cond, part){
	return res
    }
    return "next"
}

func check(cond string, part Part) bool {
    target := part[cond[0]]
    num, _ := strconv.Atoi(cond[2:])
    if cond[1] == '<'{
	return target < num
    }
    if cond[1] == '>'{
    	return target > num
    }
    panic("wrong input")
    return false
}

func getWorkflow (ops Workflow, part Part) string {
    for _, op := range(ops) {
        res := getOp(op, part)
        if res != "next" {
	    return res
        }
    }
    return ops[len(ops)-1]
}

func filter (start, end Part, op string, mode bool) (Part, Part, string) {
    split := strings.Split(op, ":")
    if len(split) == 1 {
	return start, end, op
    }
    cond := split[0]
    res := split[1]
    target := cond[0]
    num, _ := strconv.Atoi(cond[2:])
    start_copy := make(Part)
    end_copy := make(Part)
    for key := range(start) {
	start_copy[key] = start[key]
	end_copy[key] = end[key]
    }
    if mode{
        if cond[1] == '<'{
	    end_copy[target] = min(end[target], num-1)
        }
        if cond[1] == '>'{
	    start_copy[target] = max(start[target], num+1)
        }
    } else {
	if cond[1] == '<'{
	    start_copy[target] = max(start[target], num)
	}
	if cond[1] == '>'{
	    end_copy[target] = min(end[target], num)
	}
    }
    return start_copy, end_copy, res
}

func min (a, b int) int{
    if a<b{
	return a
    }
    return b
}
func max (a, b int) int{
    if a>b{
	return a
    }
    return b
}

func isValid (start, end Part) bool{
    for key, _ := range(start){
	if start[key] > end[key] {
	    return false
	}
    }
    return true
}
