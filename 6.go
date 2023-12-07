package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strconv"
	//"strings"
	//"math"
)

func main() {
    content, err := os.Open("6.in")
    if err != nil {
    	return
    }
    defer content.Close()

    scanner := bufio.NewScanner(content)

    
    var time []int
    var distance []int

    for scanner.Scan(){
	text := scanner.Text()
	rm := regexp.MustCompile(`\s+`)
	newText := rm.ReplaceAllString(text, "")
        re := regexp.MustCompile(`\d+`)
	numbers := re.FindAllStringSubmatch(newText, -1)
	if len(time)==0{
	    for i:=0;i<len(numbers);i++{
		num, _ := strconv.Atoi(numbers[i][0])
	        time = append(time, num)
	    }
	} else {
	    for i:=0;i<len(numbers);i++{
	        num, _ := strconv.Atoi(numbers[i][0])
	        distance = append(distance, num)
	    }

	}
	}
    total := 1	
    for i:=0;i<len(time);i++ {
    	total *= day6(time[i], distance[i])
    }
    fmt.Println(total)
}

func day6 (time int, distance int) int{
    fmt.Println(time)
    fmt.Println(distance)
    for i:=0;i<time;i++{
	if (time-i)*i>distance{
	    return time-2*i+1
	}
    }
    fmt.Println("Not Possible")
    return 0
}

