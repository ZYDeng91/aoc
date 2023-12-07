package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strconv"
)

func main() {
    content, err := os.Open("2_in")
    if err != nil {
    	return
    }
    defer content.Close()

    scanner := bufio.NewScanner(content)
   
    total := 0
    for scanner.Scan(){
    	total += day2(scanner.Text())
    }
    fmt.Println(total)
}

func day2 (text string) int{
    //regex_game := regexp.MustCompile("Game ([0-9]+)")
    // get capture group 1 in brackets
    //game, _ := strconv.Atoi(regex_game.FindStringSubmatch(text)[1])

    regex_colors := regexp.MustCompile(`(\d+) ([rgb])`)
    colors := regex_colors.FindAllStringSubmatch(text, -1)
   
    colors_max := make(map[string]int)
    
    for _, j := range(colors){
        num, _ := strconv.Atoi(j[1])
	if colors_max[j[2]] < num {
	    colors_max[j[2]] = num
	}
    }

    return colors_max["r"]*colors_max["g"]*colors_max["b"]
}

