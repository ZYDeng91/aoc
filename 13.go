package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    content, err := os.ReadFile("13.in")
    if err != nil {
    	return
    }

    patterns := strings.Split(string(content), "\n\n")

    mats := make([][]string, 0, len(patterns))

    for _, pattern := range(patterns){
	lines := strings.Split(pattern, "\n")
        mat := make([]string, 0, len(lines))
        for _, line := range(lines){
            mat = append(mat, line) 
        }
	if mat[len(mat)-1] == ""{
            mat = mat[0:len(mat)-1][:]
	}
   	mats = append(mats, mat) 
    }

    total := 0
    for _, mat := range(mats){
	total += day13_2(mat)
    }

    fmt.Println(total)

}

func day13 (mat []string) int {
    res := 0
    res += 100*hor(mat, -1)
    res += vert(mat, -1)
    return res
}

func day13_2 (mat []string) int {
    res := 0
    original_hor := hor(mat, -1)
    original_vert := vert(mat, -1)
    for x:=0;x<len(mat);x++{
        for y:=0;y<len(mat[0]);y++{
            repaired:=repair(mat, x, y)
	    if hor(repaired, original_hor)!=original_hor{
		res+=100*hor(repaired, original_hor)
	    }
	    if vert(repaired, original_vert)!=original_vert{
		res+=vert(repaired, original_vert)
	    }
	}
    }
    return res/2
}

func hor (mat []string, ignore int) int{
    for i:=1;i<len(mat);i++{
	if i==ignore{
	    continue
	}
	if mat[i-1] == mat[i] {
	     for j:=i+1;j<len(mat);j++{
		if 2*i-j-1<0{
		    return i
		}
		if mat[j] != mat[2*i-j-1] {
	   	    goto nextline
		}
	     }
	     return i
	}
	nextline:
    }
    return 0
}

func vert (mat []string, ignore int) int{
    mat_T := transpose(mat)
    return hor(mat_T, ignore)
}

func min (a, b int) int{
    if a<b{
	return a
    } else {
	return b
    }
}

func transpose(a []string) []string {
    newArr := make([]string, len(a[0]))
    for i := 0; i < len(a[0]); i++ {
	line := make([]byte, len(a))
        for j := 0; j < len(a); j++ {
            line[j] = a[j][i]
        }
	newArr[i] = string(line)
    }

    return newArr
}

func repair(mat []string, x, y int) []string {
    line := []rune(mat[x])
    if mat[x][y] == '#'{
	line[y] = '.'
    }
    if mat[x][y] == '.'{
	line[y] = '#'
    }
    newMat := make([]string, len(mat))
    for i, row := range(mat){
	if i==x{
	   newMat[i] = string(line)
	} else {
	newMat[i] = row
	}
    }

    return newMat
}
