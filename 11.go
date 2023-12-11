package main

import (
	"fmt"
	"os"
	"strings"
)

type loc struct {
    x int
    y int
} 

func main() {
    content, err := os.ReadFile("11.in")
    if err != nil {
    	return
    }

    lines := strings.Split(string(content), "\n")

    mat := make([]string, 0, len(lines))

    for _, line := range(lines){
   	mat = append(mat, line) 
    }

    mat = mat[0:len(mat)-1][:]

    fmt.Println(day11(mat))
}

func day11 (mat []string) int {
    var galaxies []loc
    var galaxy loc
    has_x := make(map[int]bool)
    has_y := make(map[int]bool)
    for x:=0;x<len(mat);x++{
        for y:=0;y<len(mat[x]);y++{
	    if mat[x][y] == '#' {
		galaxy.x = x
		galaxy.y = y
		has_x[x] = true
		has_y[y] = true
		galaxies = append(galaxies, galaxy)
	    }
	}
    }

/*
    var empty_x, empty_y []int

    for x:=0;x<len(mat);x++{
	if has_x[x] == false {
	    empty_x = append(empty_x, x)
	}
    }


    for y:=0;y<len(mat[0]);y++{
	if has_y[y] == false {
	    empty_y = append(empty_y, y)
	}
    }

    count := 0
    for j, i:=range(galaxies){
	count = 0
	for xx:=0;xx<len(empty_x);xx++ {
	    if xx>i.x{
		break
	    }
	    count += 1
	}
	i.x += count
	count = 0
	for yy:=0;yy<len(empty_y);yy++ {
	    if yy>i.y{
		break
	    }
	    count += 1
	}
	i.y += count
	galaxies[j] = i
    }
*/

    total := 0
    
    for i:=0;i<len(galaxies);i++{
	for j:=i+1;j<len(galaxies);j++{
	    total += getDistance(galaxies[i], galaxies[j], has_x, has_y)
	}
    }

    return total
}
/*
func getDistance (a, b loc) int{
    diff_x := 0
    diff_y := 0
    if a.x>b.x{
	diff_x += a.x-b.x
    } else {
	diff_x += b.x-a.x
    }

    if a.y>b.y{
	diff_y += a.y-b.y
    } else {
	diff_y += b.y-a.y
    }

    fmt.Println(a, b, diff_x+diff_y)

    return diff_x + diff_y

}
*/

func getDistance (a, b loc, has_x, has_y map[int]bool) int{
    diff_x := 0
    diff_y := 0
    scale := 1000000 - 1
    if a.x>b.x{
	for xx:=b.x+1;xx<a.x;xx++{
	    if !has_x[xx]{
		diff_x += scale
	    }
	}
	diff_x += a.x-b.x
    } 
    if a.x<b.x{
	for xx:=a.x+1;xx<b.x;xx++{
	    if !has_x[xx]{
		diff_x += scale
	    }
	}
	diff_x += b.x-a.x
    }

    if a.y>b.y{
	for yy:=b.y+1;yy<a.y;yy++{
	    if !has_y[yy]{
		diff_y += scale
	    }
	}
	diff_y += a.y-b.y
    } 
    if a.y<b.y{
	for yy:=a.y+1;yy<b.y;yy++{
	    if !has_y[yy]{
		diff_y += scale
	    }
	}
	diff_y += b.y-a.y
    }


    return diff_x + diff_y

}
