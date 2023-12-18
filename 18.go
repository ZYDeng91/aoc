package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"math"
)

type Line struct {
    dir byte
    dist int
}
type Loc struct {
    x int
    y int
}

func main() {
    content, err := os.Open("18.in")
    if err != nil {
    	return
    }
    defer content.Close()

    scanner := bufio.NewScanner(content)

    var lines []Line
    var colors []Line

    for scanner.Scan(){
	splits := strings.Split(scanner.Text(), " ")
	var line Line
	line.dir = splits[0][0]
	line.dist, _ = strconv.Atoi(splits[1])
	var color Line
	color.dist = hex2dec(splits[2][2:7])
	color.dir = num2dir(splits[2][7])
	lines = append(lines, line)
	colors = append(colors, color)
    }

    fmt.Println(day18(lines))
    fmt.Println(day18_2(colors))
}

func day18_2 (lines []Line) int{
    current := Loc{0, 0}
    trench := []Loc{current}
    total := 0
    border := 0
    for _, line := range(lines) {
	switch line.dir {
	    case 'U':
		current.x -= line.dist
		break
	    case 'D':
	        current.x += line.dist
	        break
	    case 'L':
	        current.y -= line.dist
	        break
	    case 'R':
	        current.y += line.dist
	        break
	}
	trench = append(trench, current)
	border += line.dist
    }
    for i:=1;i<len(trench);i++{
	total += (trench[i-1].y+trench[i].y)*(trench[i-1].x-trench[i].x)
    }

    return int(math.Abs(float64(total)/2))+border/2+1
}

func day18 (lines []Line) int{
    current := Loc{0, 0}
    offset := Loc{0, 0}
    limit := Loc{0, 0}
    trench := []Loc{current}
    for _, line := range(lines) {
        for i:=0;i<line.dist;i++ {
	    switch line.dir {
		case 'U':
		    current.x -= 1
		    break
		case 'D':
		    current.x += 1
		    break
		case 'L':
		    current.y -= 1
		    break
		case 'R':
		    current.y += 1
		    break
	    }
	    trench = append(trench, current)
	}
	if current.x<offset.x {
	    offset.x = current.x
	}
	if current.y<offset.y {
	    offset.y = current.y
	}
	if current.x>limit.x {
	    limit.x = current.x
	}
	if current.y>limit.y {
	    limit.y = current.y
	}
    }
    mat := make([][]int, limit.x-offset.x+1)
    for i:=0;i<len(mat);i++{
	emptyLine := make([]int, limit.y-offset.y+1)
	for j:=0;j<len(emptyLine);j++{
	    emptyLine[j] = 1
	}
        mat[i] = emptyLine
    }
    for _, item := range(trench){
	mat[item.x-offset.x][item.y-offset.y] = 2
    }

    lava := make([]Loc, 0)
    for x:=0;x<len(mat);x++{
	if mat[x][0] == 1{
	    mat[x][0] = 0
	    lava = append(lava, Loc{x, 0})
	}
	if mat[x][len(mat[0])-1] == 1{
	    mat[x][len(mat[0])-1] = 0
	    lava = append(lava, Loc{x, len(mat[0])-1})
	}
    }
    for y:=0;y<len(mat[0]);y++{
	if mat[0][y] == 1{
	    mat[0][y] = 0
	    lava = append(lava, Loc{0, y})
	}
	if mat[len(mat)-1][y] == 1{
	    mat[len(mat)-1][y] = 0
	    lava = append(lava, Loc{len(mat)-1, y})
        }
    }
    dirs := []Loc{Loc{0,1}, Loc{0,-1}, Loc{1,0}, Loc{-1,0}}
    for len(lava)>0 {
	newLava := make([]Loc, 0)
	for _, item := range(lava) {
	    for _, dir := range(dirs) {
		if (item.x+dir.x>0)&&(item.x+dir.x<len(mat))&&(item.y+dir.y>0)&&(item.y+dir.y<len(mat[0])) {
		    if mat[item.x+dir.x][item.y+dir.y] == 1{
			mat[item.x+dir.x][item.y+dir.y] = 0
			newLava = append(newLava, Loc{item.x+dir.x, item.y+dir.y})
		    }
		}
	    }
	}
	lava = newLava
    }
    total := 0
    for x:=0;x<len(mat);x++{
	for y:=0;y<len(mat[0]);y++{
	    if mat[x][y]!=0{
		total += 1
	    }
	}
    }

    return total
}

func hex2dec(hex string) int{
    total := 0
    for i, digit := range(hex) {
	val := 0
	if digit >= 'a' && digit <= 'f'{
	    val = int(digit - 'a') + 10
	}
	if digit >= '0' && digit <= '9'{
	    val = int(digit - '0')
	}
	total += val*int(math.Pow(16, float64(len(hex)-i-1)))
    } 
    return total
}
func num2dir(num byte) byte{
    switch num{
	case '0':
	    return 'R'
	case '1':
	    return 'D'
	case '2':
	    return 'L'
	case '3':
	    return 'U'
    }
    return '0'
}
