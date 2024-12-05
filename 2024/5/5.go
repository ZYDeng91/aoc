package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	content, err := os.Open("5.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	section := 1
	separator := "|"

	nums1 := make([][]int, 0)
	nums2 := make([][]int, 0)
		
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			section = 2
			separator = ","
			continue
		}
		nums_str := strings.Split(text, separator)
		temp := make([]int, 0)
		for _, num_str := range nums_str {
			num, _ := strconv.Atoi(num_str)
			temp = append(temp, num)		
		}
		if section == 1 {
			nums1 = append(nums1, temp)
		} else {
			nums2 = append(nums2, temp)
		}
	}
	fmt.Println(day5(nums1, nums2))
}


func day5(rules [][]int, updates [][]int) (int, int) {
	rulebook := make(map[[2]int]bool)
	res1 := 0
	res2 := 0
	for _, pair := range(rules) {
		rulebook[[2]int{pair[0], pair[1]}] = true
	}
	for _, update := range(updates) {
		for i, left := range(update) {
			if i == len(update) {
				goto good
			}
			for _, right := range(update[i+1:]) {
				if rulebook[[2]int{right, left}] {
					goto bad
				}
			}
		}
		good:
			res1 += update[(len(update)-1)/2]
			continue
		bad:

			//fmt.Println(update)
			sorted := sort(rulebook, update)
			res2 += sorted[(len(sorted)-1)/2]
	}
	return res1, res2
}

func sort(rulebook map[[2]int]bool, arr []int) []int{
	res := cp(arr)
    	var swapped bool
	for {
		swapped = false
		for i:=0;i<len(arr);i++ {
			for j:=i+1;j<len(arr);j++ {
				if rulebook[[2]int{res[j], res[i]}] {
					res[i], res[j] = res[j], res[i]
					swapped = true
					//fmt.Println(res[i], res[j])
				}
			}
		}
		if !swapped {
			break
		}
	}
	//fmt.Println(res)
	return res
}

func cp(nums []int) []int {
	res := make([]int, len(nums))
	for i, v := range nums {
		res[i] = v
	}
	return res
}
