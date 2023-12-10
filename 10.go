package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    content, err := os.ReadFile("10.in")
    if err != nil {
    	return
    }

    lines := strings.Split(string(content), "\n")

    mat := make([]string, 0, len(lines))

    for _, line := range(lines){
   	mat = append(mat, line) 
    }

    mat = mat[0:len(mat)-1][:]

    fmt.Println(day10(mat)/2)
}

func day10 (mat []string) int {
    a := 0
    b := 0
 
    dir := 'l'// 'u'
    count := 0
    for x:=0;x<len(mat);x++{
	for y:=0;y<len(mat[x]);y++{
	    if (mat[x][y] == 'S'){
		a = x
		b = y
	    }
	}
    }

    visited := make([][]int, 0, len(mat))
    for range(mat){
        emptyline := make([]int, 0, len(mat[0]))
        for range(mat[0]){
    	    emptyline = append(emptyline, 1)
        }

	visited = append(visited, emptyline)
    }

    for k:=0;k<100000;k++{
	if dir == 'u'{
	    a -= 1
	}
	if dir == 'd'{
	    a += 1
	}
	if dir == 'l'{
	    b -= 1
	}
	if dir == 'r'{
	    b += 1
	}
	visited[a][b] = 2
	count += 1
	if mat[a][b] == '.'{
	    fmt.Println(a, b, mat[a][b], dir)
	    break
	}
	if mat[a][b] == 'S'{
	    break
	}
	dir = getDir(dir, mat[a][b])
    }

    // hardcode the real pipe under S
    //mat[a][b] = 'J'

    // part 2 the pipes are leaking inbetween
    // try zoom in/upscale
    mat2 := make([][]rune, 0, 2*len(mat)-1)
    for i:=0;i<2*len(mat)-1;i++{
        emptyline := make([]rune, 0, 2*len(mat[0])-1)
	for j:=0;j<2*len(mat[0])-1;j++{
    	    emptyline = append(emptyline, '.')
        }
	mat2 = append(mat2, emptyline)
    }

    for x:=0;x<len(mat);x++{
	for y:=0;y<len(mat[x]);y++{
	    if (visited[x][y] == 2){
	        mat2[2*x][2*y] = rune(mat[x][y])
	    }
	    if (x == a) && (y == b) {
		mat2[2*x][2*y] = 'J'
	    }
	}
    }

    visited2 := make([][]int, 0, 2*len(mat)-1)
    for i:=0;i<2*len(visited)-1;i++{
        emptyline := make([]int, 0, 2*len(mat[0])-1)
        for j:=0;j<2*len(visited[0])-1;j++{
    	    emptyline = append(emptyline, 1)
        }
	visited2 = append(visited2, emptyline)
    }

    for x:=0;x<len(visited);x++{
	for y:=0;y<len(visited[x]);y++{
	    visited2[2*x][2*y] = visited[x][y]
	}
    }

    // fill in the gaps
    for x:=0;x<len(mat2);x++{
	for y:=0;y<len(mat2[x]);y++{
            if mat2[x][y] == '.'{
	    if (x>0)&&((mat2[x-1][y] == '7') || (mat2[x-1][y] == 'F') || (mat2[x-1][y] == '|')) {
		mat2[x][y] = '|'
		visited2[x][y] = 2
	    } 
	    if (y>0)&&((mat2[x][y-1] == 'L') || (mat2[x][y-1] == 'F') || (mat2[x][y-1] == '-')) {
		mat2[x][y] = '-'
		visited2[x][y] = 2
	    }
	    }
	}
    }


    // infest mode
    for x:=0;x<len(visited2);x++{
	if visited2[x][0] == 1{
	    visited2[x][0] = 0
	}
	if visited2[x][len(visited2[x])-1] == 1{
	    visited2[x][len(visited2[x])-1] = 0
	}
    }
    for x:=0;x<len(visited2[0]);x++{
	if visited2[0][x] == 1{
	    visited2[0][x] = 0
	}
	if visited2[len(visited2)-1][x] == 1{
	    visited2[len(visited2)-1][x] = 0
	}
    }

    changed := 1
    for changed>0{
	changed = 0
        for x:=0;x<len(visited2);x++{
	    for y:=0;y<len(visited2[x]);y++{
	        if visited2[x][y] == 0 {
	            for xx:=x-1; xx<=x+1; xx++ {
			for yy:=y-1; yy<=y+1; yy++ {
			    if (xx > 0) && (xx < len(visited2)) && (yy > 0) && (yy < len(visited2[x])) && (visited2[xx][yy] == 1) {
				visited2[xx][yy] = 0
				changed += 1
			    }
			}
		    }
		}
	    }
        }
	fmt.Printf("changed %d \n", changed)
    }
    // only count the original
    enclosed := 0
    for x:=0;x<len(visited);x++{
	for y:=0;y<len(visited[x]);y++{
	    if visited2[2*x][2*y] == 1 {
		enclosed += 1
	    }
	}
    }
    fmt.Println(enclosed)

    return count
}

func getDir(last rune, current byte) rune{
    if last == 'u'{
	if current == '|' {
	    return 'u'
	}
	if current == '7' {
	    return 'l'
	}
	if current == 'F' {
	    return 'r'
	}
    }

    if last == 'd'{
	if current == '|' {
	    return 'd'
	}
	if current == 'J' {
	    return 'l'
	}
	if current == 'L' {
	    return 'r'
	}
    }

    if last == 'l'{
	if current == '-' {
	    return 'l'
	}
	if current == 'F' {
	    return 'd'
	}
	if current == 'L' {
	    return 'u'
	}
    }

    if last == 'r'{
	if current == '-' {
	    return 'r'
	}
	if current == '7' {
	    return 'd'
	}
	if current == 'J' {
	    return 'u'
	}
    }
    fmt.Println("out of pipe")
    return '0'

}

