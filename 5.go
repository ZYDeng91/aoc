package main

import (
	"fmt"
	"os"
	"bufio"
	//"regexp"
	"strconv"
	"strings"
	//"math"
)

type mapp struct {
    dest []int
    src []int
    rang []int
}

func main() {
    content, err := os.Open("5.in")
    if err != nil {
    	return
    }
    defer content.Close()

    scanner := bufio.NewScanner(content)
    pipes := make([]mapp,7)
    index := -1
    var seeds []string

    for scanner.Scan(){
	text := scanner.Text()
	
	if text == "" {
	    continue
	}
	if text[:6] == "seeds:" {
	    seeds = strings.Split(text[7:], " ")
	}
	if text[len(text)-4:] == "map:" {
    	    index += 1
	}
   	if text[0]>='0' && text[0]<='9' {
	    numbers := strings.Split(text, " ")
   	    dest, _ := strconv.Atoi(numbers[0])
	    src, _ := strconv.Atoi(numbers[1])
	    rang, _ := strconv.Atoi(numbers[2])
	    pipes[index].dest = append(pipes[index].dest, dest) 
	    pipes[index].src = append(pipes[index].src, src)
	    pipes[index].rang = append(pipes[index].rang, rang)  
	}
	    
    }

    var seeds2 []int
    /*for i:=0;i<len(seeds)/2;i++ {
	a, _ := strconv.Atoi(seeds[2*i])
	b, _ := strconv.Atoi(seeds[2*i+1])
	for j:=a;j<a+b;j+=1000{
	    seeds2 = append(seeds2, j)
	}
    }*/
    fmt.Println(len(seeds))
    for j:=4246766536;j<4246806536;j+=1{
	    seeds2 = append(seeds2, j)
    }
    fmt.Println(len(seeds2))
    
    min := -1
    for _, seed := range(seeds2) {
	//input, _ := strconv.Atoi(seed)
	input := seed
	for _, pipe := range(pipes) {
	    input = day5(input, pipe)
	}
	if (min==-1)||(min>input) {
	    min = input
	    fmt.Println(seed,input)
	}
    }

}

func day5 (input int, pipe mapp) int{
    for i:=0;i<len(pipe.dest);i++{
	if (input>=pipe.src[i] && input<pipe.src[i]+pipe.rang[i]){
	    return input-pipe.src[i]+pipe.dest[i]
	}
    }
    return input
}

