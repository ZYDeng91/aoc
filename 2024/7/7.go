package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	content, err := os.Open("7.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	total := 0
	total2 := 0
	for scanner.Scan() {
		text := scanner.Text()
		nums_str := strings.Split(text, " ")
		nums := make([]int, 0)
		left := 0
		for i, num_str := range nums_str {
			if i == 0 {
				left, _ = strconv.Atoi(num_str[:len(num_str)-1])
				continue
			}
			num, _ := strconv.Atoi(num_str)
			nums = append(nums, num)
		}
		tmp := day7(left, nums, left)
		total += tmp
		if tmp == 0 {
			total2 += day7p2(left, nums, left)
		}
	}
	fmt.Println(total, total+total2)
}

func day7(left int, right []int, total int) int {
	if len(right) == 1 {
		if right[0] == left {
			return total
		}
		return 0
	}
	last := len(right)-1;
	if left % right[last] == 0 {
		temp1 := day7(left/right[last], right[:last], total) 
		temp2 := day7(left-right[last], right[:last], total)
		if temp1 != temp2 {
			return temp1+temp2
		}
		return temp1
	} else {
		return day7(left-right[last], right[:last], total)
	}
}

func day7p2(left int, right []int, total int) int {
	if len(right) == 1 {
		if right[0] == left {
			return total
		}
		return 0
	}
	last := len(right)-1;
	if unconcat(left, right[last]) != -1 {
		temp1 := day7p2(unconcat(left, right[last]), right[:last], total)
		temp2 := 0
		if left % right[last] == 0 {
			temp21 := day7p2(left/right[last], right[:last], total) 
			temp22 := day7p2(left-right[last], right[:last], total)
			if temp21 != temp22 {
				temp2 = temp21+temp22
			} else {
				temp2 = temp21
			}
		} else {
			temp2 = day7p2(left-right[last], right[:last], total)
		}
		if temp1 != temp2 {
			return temp1+temp2
		}
		return temp1

	} else if left % right[last] == 0 {
		temp1 := day7p2(left/right[last], right[:last], total) 
		temp2 := day7p2(left-right[last], right[:last], total)
		if temp1 != temp2 {
			return temp1+temp2
		}
		return temp1
	}
	return day7p2(left-right[last], right[:last], total)

}

func unconcat(a, b int) int {
	if a <= b {
		return -1
	}
	len_b := 1
	for len_b <= b {
		len_b *= 10
	}
	if (a - b)%len_b != 0 {
		return -1
	}
	return (a - b)/len_b
}
