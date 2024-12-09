package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
	"strconv"
)

type File struct {
	start int
	size int
	id int
}

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


func day9(files []int, spaces []int) int {
	files2 := make([]File, 0)
	spaces2 := make([]File, 0)
	ptr := 0
	for i:=0;i<len(files)-1;i++ {
		files2 = append(files2, File{start: ptr, size: files[i], id: i})
		ptr += files[i]
		spaces2 = append(spaces2, File{start: ptr, size: spaces[i], id: -1})
		ptr += spaces[i]
	}
	files2 = append(files2, File{start: ptr, size: files[len(files)-1], id: len(files)-1})
	for f:=len(files2)-1;f>=0;f-- {
		for ss:=0;ss<len(spaces2);ss++ {
			s := spaces2[ss]
			if files2[f].start < s.start {
				break
			}
			if files2[f].size <= s.size {
				files2[f].start = s.start
				s.start += files2[f].size
				s.size -= files2[f].size
				spaces2[ss] = s
				break
			}
		}
	}
	return countChecksum(files2)
}

func countChecksum(files []File) int {
	res := 0
	for _, f := range(files) {
		res += checksum(f)
	}
	return res
}

func checksum(f File) int {
	return f.id * (2*f.start + f.size - 1) * f.size / 2
}

func count(arr []int) int {
	res := 0
	for _, v := range(arr) {
		res += v
	}
	return res
}
