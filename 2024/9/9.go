package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
	"strconv"
)

func main() {
	content, err := os.Open("9.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	nums1 := make([]int, 0)
	nums2 := make([]int, 0)
		
	for scanner.Scan() {
		text := scanner.Text()
		for i, v := range(text) {
			num, _ := strconv.Atoi(string(v))
			if i % 2 == 0 {
				nums1 = append(nums1, num)
			} else {
				nums2 = append(nums2, num)
			}
		}
	}
	fmt.Println(day9(nums1, nums2))
}


func day9(files []int, spaces []int) (int, int) {
	res1 := 0
	res2 := 0
	files_total := count(files)
	files_left := files_total
	spaces_count := 0
	i := 0
	for i=0;i<len(files);i++ {
		spaces_count += spaces[i]
		files_left -= files[i]
		if spaces_count > files_left {
			break
		}
	}
	//fmt.Println(files_total,spaces_count)

	file_loc := make([]int, files_total - files_left)
	space_loc := make([]int, spaces_count)
	file_id := make([]int, files_total - files_left)
	space_id := make([]int, spaces_count)

	f_size := 0
	s_size := 0
	leftover := 0

	out:
	for j:=0;j<=i;j++ {
		for f:=0;f<files[j];f++ {
			if f_size + s_size + f >= files_total {
				f_size += f
				leftover = files[j] - f 
				break out
			}
			file_loc[f_size+f] = f_size + s_size + f
			file_id[f_size+f] = j
		}
		f_size += files[j]
		for s:=0;s<spaces[j];s++ {
			space_loc[s_size+s] = f_size + s_size + s
		}
		s_size += spaces[j]
	}
	//fmt.Println(file_loc, file_id)
	s_size = 0
	for j:=len(files)-1;j>i;j-- {
		for s:=0;s<files[j];s++{
			space_id[s_size+s] = j
		}
		s_size += files[j]
	}
	for s:=0;s<leftover;s++ {
		space_id[s_size+s] = i
	}
	//fmt.Println(space_loc, space_id)
	res1 += arr_mul(file_loc, file_id)
	res1 += arr_mul(space_loc, space_id)
	return res1, res2
}

func arr_mul(arr1, arr2 []int) int {
	res := 0
	for i:=0;i<len(arr1);i++ {
		res += arr1[i]*arr2[i]
	}
	return res
}

func count(arr []int) int {
	res := 0
	for _, v := range(arr) {
		res += v
	}
	return res
}

func cp(nums []int) []int {
	res := make([]int, len(nums))
	for i, v := range nums {
		res[i] = v
	}
	return res
}
