package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Module struct {
    Name string
    Prefix byte
    Receiver []string
    // 0: recently sent a low
    // 1: recently sent a high
    State int
}

func main() {
    content, err := os.Open("20.in")
    if err != nil {
    	return
    }
    defer content.Close()

    scanner := bufio.NewScanner(content)

    lines := make([]Module, 0)


    for scanner.Scan(){
	text := scanner.Text()
	splits := strings.Split(text, " -> ")
	var line Module
	line.Name = splits[0][1:]
	line.Prefix = splits[0][0]
	line.Receiver = strings.Split(splits[1], ", ")
	line.State = 0
	lines = append(lines, line)
    }
    //fmt.Println(day20(lines, 1000)
    fmt.Println(day20(lines, 100000))
}


func day20 (lines []Module, button int) int{
    low := 0 
    high := 0
    modules := make(map[string]Module)
    for _, line := range(lines) {
	modules[line.Name] = line
    }

    watchers := make(map[string][]string)
    for _, line := range(lines) {
	for _, receiver := range(line.Receiver) {
	    if modules[receiver].Prefix == '&' {
		watchers[receiver] = append(watchers[receiver], line.Name)
	    }
	}

    }
    for i:=0;i<button;i++{
    low += 1
    pulses := []string{"roadcaster"}

    for len(pulses)>0{
	newPulses := make([]string, 0)
	for _, pulse := range(pulses) {
	    current := modules[pulse]
	    // 'b'roadcaster start with state 0, it can directly proceed
	    if current.Prefix == '%' {
		current.State = 1 - current.State
	    }
	    if current.Prefix == '&' {
		current.State = 0
		for _, watched := range(watchers[pulse]) {
		    if modules[watched].State == 0 {
			current.State = 1
			break
		    }
		}
	    }
            if current.State == 0 {
		low += len(current.Receiver)
	    }
	    if current.State == 1 {
		high += len(current.Receiver)
	    }
	    for _, receiver := range(current.Receiver) {
		if modules[receiver].Prefix == '%' && current.State == 1 {
		    continue
		}
		// &sq -> rx 
		// &fv, &kk, &vt, &xr -> sq
		// &vm -> fv, &kb -> kk, &dn -> vt, &vk -> xr
		// we get rs when all of fv, kk, vt, xr are high
		// find their period of change, then find lcm
		// 3863, 3931, 3797, 3769
		/*if receiver == "xr" && current.State == 0 {
		    fmt.Println(i, current.Name)
		}*/
		
		if receiver == "rx" && current.State == 0 {
		    fmt.Println(i)
		}
	        newPulses = append(newPulses, receiver)
	    }
	    modules[pulse] = current
	}
	pulses = newPulses
    }
    }
    fmt.Println(low, high)

    return low*high
}
