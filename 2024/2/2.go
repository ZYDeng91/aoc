package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	content, err := os.Open("2.in")
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
		for _, num_str := range nums_str {
			num, _ := strconv.Atoi(num_str)
			nums = append(nums, num)
		}
		total += day2(nums)
		total2 += day2p2(nums)
	}
	fmt.Println(total)
	fmt.Println(total2)
}

func day2(nums []int) int {
	d1 := nums[1]-nums[0]
	for i:=0;i<len(nums)-1;i++{
		d := nums[i+1]-nums[i]
		if d*d1 <= 0 || d > 3 || d < -3 {
			return 0
		}
	}
	return 1
}

func day2p2(nums []int) int {
	d1 := nums[1]-nums[0]
	pos := -1
	for i:=0;i<len(nums)-1;i++{
		d := nums[i+1]-nums[i]
		if d*d1 <= 0 || d > 3 || d < -3 {
			pos = i
			break
		}
	}
	if pos == -1 {
		return 1
	}
	if pos == 1 {
		if day2(cp(nums[1:])) == 1 {
			return 1
		}
	}

	// remove left
	nums1 := append(cp(nums[:pos]), nums[pos+1:]...)
	// remove right
	nums2 := append(cp(nums[:pos+1]), nums[pos+2:]...)

	if day2(nums1) == 1 || day2(nums2) == 1 {
		return 1
	}

	return 0
}

func cp(nums []int) []int {
	res := make([]int, len(nums))
	for i, v := range nums {
		res[i] = v
	}
	return res
}
