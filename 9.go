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

func main() {
    content, err := os.Open("9.in")
    if err != nil {
    	return
    }
    defer content.Close()

    scanner := bufio.NewScanner(content)

    total := 0
    for scanner.Scan(){
	total += day9(scanner.Text())
    }
    fmt.Println(total)
}

func day9 (text string) int {
    numbers_str := strings.Split(text, " ")
    var numbers []int
    for _, number_str := range(numbers_str) {
	number, _ := strconv.Atoi(number_str)
	numbers = append(numbers, number)
    }

    var numbers_all [][]int
    numbers_all = append(numbers_all, numbers)

    extra := -10000
    for a:=0;a<1000;a++{
	numbers_all = append(numbers_all, getNext(numbers_all[len(numbers_all)-1]))
	if isLast(numbers_all[len(numbers_all)-1]){
	    extra = numbers_all[len(numbers_all)-1][0]
	    break
	}
    }

    if extra == -10000 {
	fmt.Println("failed to diverge after 1k derivatives")
	return 0
    } else {
	/*
	for i:=len(numbers_all)-2;i>=0;i-- {
	    extra += numbers_all[i][len(numbers_all[i])-1]
	}*/
        for i:=len(numbers_all)-2;i>=0;i-- {
	    extra = numbers_all[i][0] - extra
	}
    }
    return extra

}

func getNext(numbers []int) []int {
    var result []int
    for i:=1; i<len(numbers); i++ {
	result = append(result, numbers[i]-numbers[i-1])
    }

    return result
}

func isLast(numbers []int) bool {
    for i:=1;i<len(numbers);i++{
	if numbers[i] != numbers[i-1]{
	    return false
	}
    }
    return true
}
