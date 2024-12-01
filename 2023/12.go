package main

import (
	"bufio"
	"fmt"
	"os"
	//"regexp"
	"strconv"
	"strings"
	//"math"
)

type pair struct {
	i int
	j int
}

func main() {
	content, err := os.Open("12.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	total := 0
	for scanner.Scan() {
		total += day12(scanner.Text())
	}
	fmt.Println(total)
}

func day12(text string) int {
	record := strings.Split(text, " ")[0]
	//re := regexp.MustCompile(`[\?#]+`)
	//locs := re.FindAllStringSubmatch(left, -1)

	groups_str := strings.Split(strings.Split(text, " ")[1], ",")
	var groups []int
	for _, group_str := range groups_str {
		group, _ := strconv.Atoi(group_str)
		groups = append(groups, group)
	}

	record5x := []string{record, record, record, record, record}
	record5 := strings.Join(record5x[:], "?")

	var groups5 []int
	for i := 0; i < 5; i++ {
		groups5 = append(groups5, groups[:]...)
	}

	cache := make(map[pair]int)

	return dp(0, 0, record5, groups5, cache)
}

func dp(i, j int, record string, groups []int, cache map[pair]int) int {
	if i >= len(record) {
		if j < len(groups) {
			return 0
		}
		return 1
	}

	// use cache if possible (memoization)
	if val, ok := cache[pair{i, j}]; ok {
		return val
	}

	// continue
	if record[i] == '.' {
		return dp(i+1, j, record, groups, cache)
	}

	res := 0
	// start as current
	if j < len(groups) {
		count := 0
		for count = 0; i+count < len(record); count++ {
			// too many '#'s     || ended by '.'           || end on '?' if length is exactly enough
			if count > groups[j] || record[i+count] == '.' || (record[i+count] == '?' && count == groups[j]) {
				break
			}
		}
		// proceed to the next group
		if count == groups[j] {
			res += dp(i+count+1, j+1, record, groups, cache)
		}
	}

	// current + continue => ? branches as both # and .
	if record[i] == '?' {
		res += dp(i+1, j, record, groups, cache)
	}

	cache[pair{i, j}] = res
	return res
}

/*
func getArr(loc string, group int) int{
    a := -1
    b := -1
    for i:=0;i<len(loc);i++{
	if (a==-1) && (loc[i]=='#'){
	    a = i
	}
	if loc[i]=='#'{
	    b = i
	}
    }
    if (group>len(loc))||(b-a>group) {
	fmt.Println("invalid split")
	return 0
    }
    if group == 0{
	return 1
    }

    if a==-1{
	return len(loc)-group+1
    } else {
	core := b-a
	free := len(loc)-core
        left := min(a, free)
	right := min(len(loc)-1-b, free)
	return left+right-free+1
    }

}

func min(a, b int) int{
    if a<b{
	return a
    }
    return b
}*/
