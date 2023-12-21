package main

import (
	"fmt"
	"os"
	"strings"
)

type Loc struct {
    x int
    y int
}

type Meta struct {
    meta Loc
    loc Loc
}

func main() {
    content, err := os.ReadFile("21.in")
    if err != nil {
    	return
    }

    lines := strings.Split(string(content), "\n")

    mat := make([]string, len(lines))

    for i, line := range(lines){
   	mat[i] = line
    }

    mat = mat[0:len(mat)-1][:]

    fmt.Println(day21(mat))
    fmt.Println(day21_2(mat))
}


func day21_2 (mat []string) int {
    var start Meta
    start.meta = Loc{0, 0}
    for x:=0;x<len(mat);x++{
	for y:=0;y<len(mat[0]);y++{
	    if mat[x][y] == 'S' {
		start.loc = Loc{x,y}
	    }
	}
    }

    current := []Meta{start}
    total := 0
    steps := 2001

    totals := make([]int, steps+1)
    totals[0] = 1
    
    seen := make(map[Meta]bool)
    for i:=1;i<=steps;i++{
	current = step_2(mat, current)
	next := make([]Meta, 0)
	for _, item := range(current) {
	    if !seen[item]{
	        next = append(next, item)
		seen[item] = true
		_, oddity := divide((item.meta.x*len(mat)+item.loc.x-start.loc.x)+(item.meta.y*len(mat[0])+item.loc.y-start.loc.y),2)
	        if oddity%2 == steps%2 {
		    total += 1
		}
	    }
	}
	current = next
	totals[i] = total
	//fmt.Println(i, totals[i], totals[i]-totals[i-1])
    }

    // the result is a quadratic equation
    // think as if the map has no obstacles

    // adding obstacles would be only adding extra steps
    // there should be a pattern revolving around the width of the matrix (which is a square)
    // if only counting odd/even steps, check two times that width

    // in the example every 22 steps, new available plots: totals[i]-totals[i-1] += 50
    
    // keep taking derivatives

    // 26501365%(2*len(mat)) = 101150, 65
    // a0 = 90676, a(n)-a(n-1) = 90676 + n*120720


    i := 2*len(mat)+65
    m := 0
    res := make([]int, 101151)
    res[0] = totals[65]
    for n:=1;n<101151;n++{
	res[n] = res[n-1] + 90676 + (n-1)*120720
    }
    for i<len(totals){
	/*fmt.Println(i, totals[i]-totals[i-1], totals[i-len(mat)]-totals[i-len(mat)-1])*/
	fmt.Println(i, totals[i], totals[i]-totals[i-2*len(mat)])
	fmt.Println(m, res[m+1], 90676+m*120720)
	i+=2*len(mat)
	m += 1
    }


    //return total
    return res[101150]
}

func step_2 (mat []string, current []Meta) []Meta{
    res := make([]Meta, 0)
    for _, item := range(current){
	for _, dir := range([]Loc{Loc{0,1},Loc{0,-1},Loc{1,0},Loc{-1,0}}){
	    newmetax, newx := divide(item.loc.x+dir.x, len(mat))
	    newmetay, newy := divide(item.loc.y+dir.y, len(mat[0]))
	    if mat[newx][newy] == '#'{
		continue	
	    }
	    res = append(res, Meta{Loc{item.meta.x+newmetax, item.meta.y+newmetay}, Loc{newx, newy}})
	}
    }
    return res
}

func divide (a, b int) (int, int){
    if a < 0 {
	return a/b-1, b+a%b
    }
    return a/b, a%b
}


func day21 (mat []string) int {
    var start Loc
    for x:=0;x<len(mat);x++{
	for y:=0;y<len(mat[0]);y++{
	    if mat[x][y] == 'S' {
		start.x = x
		start.y = y
	    }	
	}
    }
    current := []Loc{start}
    total := 0
    visited := getEmptyMat(mat)
    steps := 6
    for i:=1;i<=steps;i++{
	current = step(mat, current)
	next := make([]Loc, 0)
	for _, item := range(current) {
	    if visited[item.x][item.y] == 0{
	        visited[item.x][item.y] = i%2+1
		next = append(next, item)
	    }
	}
	current = next
    }
    for x:=0;x<len(visited);x++{
	for y:=0;y<len(visited[0]);y++{
	    if visited[x][y] == steps%2+1 {
	   	total += 1 
	    }	
	}
    }

    return total
}

func step (mat []string, current []Loc) []Loc{
    res := make([]Loc, 0)
    for _, item := range(current){
	for _, dir := range([]Loc{Loc{0,1},Loc{0,-1},Loc{1,0},Loc{-1,0}}){
	    newx := item.x+dir.x
	    newy := item.y+dir.y
	    if newx<0||newy<0||newx>=len(mat)||newy>=len(mat[0]) {
		continue
	    }
	    if mat[newx][newy] == '#'{
		continue	
	    }
	    res = append(res, Loc{newx, newy})
	}
    }
    return res
}


func getEmptyMat (mat []string) [][]int {
    visited := make([][]int, 0, len(mat))
    for range(mat){
        emptyline := make([]int, 0, len(mat[0]))
        for range(mat[0]){
    	    emptyline = append(emptyline, 0)
        }
	visited = append(visited, emptyline)
    }
    return visited
}

func sum (input []int) int {
    res := 0
    for _, val := range(input){
	res += val
    }
    return res
}
